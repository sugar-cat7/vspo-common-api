package ports

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/sugar-cat7/vspo-common-api/domain/entities"
	"github.com/sugar-cat7/vspo-common-api/domain/ports"
	"github.com/sugar-cat7/vspo-common-api/util"
)

type discordServiceImpl struct {
	Session *discordgo.Session
}

func NewDiscordService() (ports.DiscordService, error) {
	botToken, ok := os.LookupEnv("DISCORD_BOT_TOKEN")
	if !ok {
		return nil, fmt.Errorf("DISCORD_BOT_TOKEN not set")
	}

	session, err := discordgo.New("Bot " + botToken)
	if err != nil {
		return nil, err
	}

	return &discordServiceImpl{Session: session}, nil
}

func (s *discordServiceImpl) SendMessages(liveStreams entities.Videos, countryCode string) error {
	botUser, err := s.Session.User("@me")
	if err != nil {
		return fmt.Errorf("error getting bot user: %v", err)
	}

	var guilds []*discordgo.UserGuild
	var lastID string
	for {
		g, err := s.Session.UserGuilds(200, "", lastID)
		if err != nil {
			return fmt.Errorf("error getting user guilds: %v", err)
		}
		if len(g) == 0 {
			break
		}
		guilds = append(guilds, g...)
		lastID = g[len(g)-1].ID

		time.Sleep(1 * time.Second)
	}

	var errs []string
	const batchSize = 50
	for i := 0; i < len(guilds); i += batchSize {
		end := i + batchSize
		if end > len(guilds) {
			end = len(guilds)
		}
		batch := guilds[i:end]

		var wg sync.WaitGroup
		errCh := make(chan error, len(batch))

		for _, guild := range batch {
			wg.Add(1)
			go func(guild *discordgo.UserGuild) {
				defer wg.Done()
				if err := s.processGuild(guild, liveStreams, botUser, countryCode); err != nil {
					errCh <- fmt.Errorf("error processing guild %s: %v", guild.Name, err)
				}
			}(guild)
		}

		wg.Wait()
		close(errCh)

		for err := range errCh {
			errs = append(errs, err.Error())
		}

		// FIXME: temp...Sleep between batches to avoid overwhelming the server
		time.Sleep(4 * time.Second)
	}
	if len(errs) > 0 {
		return fmt.Errorf("errors processing guilds: %s", strings.Join(errs, "; "))
	}
	return nil
}

func (s *discordServiceImpl) processGuild(guild *discordgo.UserGuild, liveStreams entities.Videos, botUser *discordgo.User, countryCode string) error {
	targetChannelName := "ぶいすぽ配信情報"
	initialMessage := "すぽじゅーるは、ぶいすぽっ!メンバーの配信(Youtube/Twitch/ツイキャス/ニコニコ)や切り抜きを一覧で確認できる非公式サイトです。 /Spodule aggregates schedules for Japan's Vtuber group, Vspo.\n\nWeb版はこちら：https://www.vspo-schedule.com/schedule/all"

	channels, err := s.Session.GuildChannels(guild.ID)
	if err != nil {
		return fmt.Errorf("error getting guild channels: %v", err)
	}

	var targetChannel *discordgo.Channel
	for _, channel := range channels {
		if channel.Name == targetChannelName {
			targetChannel = channel
			break
		}
	}

	if targetChannel == nil {
		targetChannel, err = s.Session.GuildChannelCreate(guild.ID, targetChannelName, discordgo.ChannelTypeGuildText)
		if err != nil {
			return fmt.Errorf("error creating channel %s: %v", targetChannelName, err)
		}
	}
	// 新しいライブストリームの一覧を元に埋め込みメッセージを作成
	newEmbeds, newEmbedMap := buildEmbeds(liveStreams, countryCode)

	// 既存のメッセージの取得
	existingMessages, err := s.Session.ChannelMessages(targetChannel.ID, 100, "", "", "")
	if err != nil {
		return fmt.Errorf("error getting messages from channel %s: %v", targetChannel.Name, err)
	}

	// ターゲットチャンネルにメッセージが一件も投稿されていない場合、initmessageを投稿します。
	if len(existingMessages) == 0 {
		_, err = s.Session.ChannelMessageSend(targetChannel.ID, initialMessage)
		if err != nil {
			return fmt.Errorf("error sending initial message to channel %s: %v", targetChannel.Name, err)
		}
	}

	existingEmbeds := make(map[string]*discordgo.MessageEmbed)
	for _, msg := range existingMessages {
		for _, embed := range msg.Embeds {
			existingEmbeds[embed.URL] = embed
		}
	}

	newMessageEmbedMap := make(map[string]*discordgo.MessageEmbed)
	for _, embed := range newEmbeds {
		newMessageEmbedMap[embed.URL] = embed
	}

	for _, newEmbed := range newEmbeds {
		_, exists := existingEmbeds[newEmbed.URL]
		if !exists {
			// If not exists, send the new embed message
			_, err := s.Session.ChannelMessageSendComplex(targetChannel.ID, &discordgo.MessageSend{
				Embed: newEmbed,
			})
			if err != nil {
				return fmt.Errorf("error sending embed message to channel %s: %v", targetChannel.Name, err)
			}

		}
	}

	for _, msg := range existingMessages {
		if msg.Content == initialMessage {
			continue
		}
		for _, embed := range msg.Embeds {
			if _, exists := newEmbedMap[embed.URL]; !exists {
				// _, err := s.Session.ChannelMessage(targetChannel.ID, embed.URL)
				// if err == nil {
				if msg.Content != initialMessage && msg.Author.ID == botUser.ID {
					// 既存のメッセージの中で新しいライブストリームの一覧にないものがあれば、それを削除
					err := s.Session.ChannelMessageDelete(targetChannel.ID, msg.ID)
					if err != nil {
						return fmt.Errorf("error deleting message in channel %s: %v", targetChannel.Name, err)
					}
					// }
				}
			}
			reSendURL := make([]string, 0, len(newEmbeds))

			t, err := util.ParseTimeForCountry(embed.Fields[0].Value, countryCode)
			if err != nil {
				return err
			}
			isFuture, err := util.IsFuture(t, countryCode)
			if err != nil {
				return err
			}
			// ラベルの色の修正
			if !isFuture && embed.Color != entities.ColorLive {
				if msg.Content != initialMessage && msg.Author.ID == botUser.ID {
					// 過去の配信は削除
					err := s.Session.ChannelMessageDelete(targetChannel.ID, msg.ID)
					reSendURL = append(reSendURL, embed.URL)
					if err != nil {
						return fmt.Errorf("error deleting message in channel %s: %v", targetChannel.Name, err)
					}
				}
			}
			// twitchの配信が終了したら削除（配信URLが変わるため別途判定が必要?）
			// if strings.Contains(embed.Image.URL, "static-cdn.jtvnw.net") {
			// 	err := s.Session.ChannelMessageDelete(targetChannel.ID, msg.ID)
			// 	if err != nil {
			// 		return fmt.Errorf("error deleting message in channel %s: %v", targetChannel.Name, err)
			// 	}
			// }

			if len(reSendURL) > 0 {
				// 削除したメッセージを再送信
				for _, url := range reSendURL {
					if _, ok := newMessageEmbedMap[url]; !ok {
						_, err := s.Session.ChannelMessageSendComplex(targetChannel.ID, &discordgo.MessageSend{
							Embed: newMessageEmbedMap[url],
						})
						if err != nil {
							return fmt.Errorf("error sending embed message to channel %s: %v", targetChannel.Name, err)
						}
					}
				}
			}
		}
	}

	return nil
}

func buildEmbeds(liveStreams entities.Videos, countryCode string) ([]*discordgo.MessageEmbed, map[string]bool) {
	var embeds []*discordgo.MessageEmbed
	isExistVideoMap := make(map[string]bool)

	for _, video := range liveStreams {
		formattedTime, err := util.FormatTimeForCountry(video.ScheduledStartTime, countryCode)
		if err != nil {
			continue
		}
		embed := &discordgo.MessageEmbed{
			Title: video.Title,
			URL:   video.Link,
			Color: video.GetStatusColor(),
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "配信日時",
					Value:  formattedTime,
					Inline: true,
				},
			},
			Image: &discordgo.MessageEmbedImage{
				URL: video.Thumbnails.Default.URL,
			},
			Author: &discordgo.MessageEmbedAuthor{
				Name:    video.ChannelTitle,
				IconURL: video.ChannelIcon,
			},
			Footer: &discordgo.MessageEmbedFooter{
				Text:    video.Platform.Upper() + " " + " Powered by すぽじゅーる",
				IconURL: video.Platform.GetPlatformIconURL(),
			},
		}
		embeds = append(embeds, embed)

		isExistVideoMap[video.Link] = true

	}

	return embeds, isExistVideoMap
}

func (s *discordServiceImpl) DeleteMessages() error {
	var guilds []*discordgo.UserGuild
	var lastID string
	for {
		g, err := s.Session.UserGuilds(200, "", lastID)
		if err != nil {
			return fmt.Errorf("error getting user guilds: %v", err)
		}
		if len(g) == 0 {
			break
		}
		guilds = append(guilds, g...)
		lastID = g[len(g)-1].ID

		time.Sleep(1 * time.Second)
	}

	var errs []string
	const batchSize = 50
	for i := 0; i < len(guilds); i += batchSize {
		end := i + batchSize
		if end > len(guilds) {
			end = len(guilds)
		}
		batch := guilds[i:end]

		var wg sync.WaitGroup
		errCh := make(chan error, len(batch))

		for _, guild := range batch {
			wg.Add(1)
			go func(guild *discordgo.UserGuild) {
				defer wg.Done()
				if err := s.DeleteAllExceptInitFromGuild(guild.ID); err != nil {
					errCh <- fmt.Errorf("error processing guild %s: %v", guild.Name, err)
				}
			}(guild)
		}

		wg.Wait()
		close(errCh)

		for err := range errCh {
			errs = append(errs, err.Error())
		}

		// FIXME: temp...Sleep between batches to avoid overwhelming the server
		time.Sleep(4 * time.Second)
	}
	if len(errs) > 0 {
		return fmt.Errorf("errors processing guilds: %s", strings.Join(errs, "; "))
	}
	return nil
}

func (s *discordServiceImpl) DeleteAllExceptInitFromGuild(guildID string) error {
	botUser, err := s.Session.User("@me")
	if err != nil {
		return fmt.Errorf("error getting bot user: %v", err)
	}
	const initialMessage = "すぽじゅーるは、ぶいすぽっ!メンバーの配信(Youtube/Twitch/ツイキャス/ニコニコ)や切り抜きを一覧で確認できる非公式サイトです。 /Spodule aggregates schedules for Japan's Vtuber group, Vspo.\n\nWeb版はこちら：https://www.vspo-schedule.com/schedule/all"
	targetChannelName := "ぶいすぽ配信情報"
	// Fetch all channels in the guild
	channels, err := s.Session.GuildChannels(guildID)
	if err != nil {
		return fmt.Errorf("error getting guild channels for guild %s: %v", guildID, err)
	}
	var targetChannel *discordgo.Channel
	for _, channel := range channels {
		if channel.Name == targetChannelName {
			targetChannel = channel
			break
		}
	}

	if targetChannel != nil {
		// Fetch existing messages from the channel
		existingMessages, err := s.Session.ChannelMessages(targetChannel.ID, 100, "", "", "")
		if err != nil {
			return fmt.Errorf("error getting messages from channel ID %s: %v", targetChannel.ID, err)
		}

		// Iterate over each message
		for _, msg := range existingMessages {
			// If the message content isn't the initialMessage, delete it
			if msg.Content != initialMessage && msg.Author.ID == botUser.ID {
				err := s.Session.ChannelMessageDelete(targetChannel.ID, msg.ID)
				if err != nil {
					return fmt.Errorf("error deleting message with ID %s in channel ID %s: %v", msg.ID, targetChannel.ID, err)
				}
			}
		}
	}

	return nil
}
