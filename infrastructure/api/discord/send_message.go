package ports

import (
	"fmt"
	"os"
	"strings"
	"sync"

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

func (s *discordServiceImpl) SendMessages(liveStreams []*entities.Video, countryCode string) error {
	botUser, err := s.Session.User("@me")
	if err != nil {
		return fmt.Errorf("error getting bot user: %v", err)
	}

	// 所属する全てのサーバー（ギルド）の取得
	guilds, err := s.Session.UserGuilds(100, "", "")
	if err != nil {
		return fmt.Errorf("error getting user guilds: %v", err)
	}

	var wg sync.WaitGroup
	errCh := make(chan error, len(guilds))

	for _, guild := range guilds {
		wg.Add(1)
		go func(guild *discordgo.UserGuild) {
			defer wg.Done()
			if err := s.processGuild(guild, liveStreams, botUser, countryCode); err != nil {
				errCh <- fmt.Errorf("error processing guild %s: %v", guild.Name, err)
			}
		}(guild)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	close(errCh)

	// Collect all errors, if any
	var errs []string
	for err := range errCh {
		errs = append(errs, err.Error())
	}

	if len(errs) > 0 {
		return fmt.Errorf("errors processing guilds: %s", strings.Join(errs, "; "))
	}

	return nil
}

func (s *discordServiceImpl) processGuild(guild *discordgo.UserGuild, liveStreams []*entities.Video, botUser *discordgo.User, countryCode string) error {
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

	for _, newEmbed := range newEmbeds {
		oldEmbed, exists := existingEmbeds[newEmbed.URL]
		if exists {
			// Check if there is any change, for simplicity, we will check formattedTime and Image only
			// You can add more fields as needed.
			if oldEmbed.Fields[0].Value != newEmbed.Fields[0].Value || oldEmbed.Image.URL != newEmbed.Image.URL || oldEmbed.Color != newEmbed.Color {
				_, err := s.Session.ChannelMessage(targetChannel.ID, oldEmbed.URL)
				if err == nil {
					_, err := s.Session.ChannelMessageEditComplex(&discordgo.MessageEdit{
						ID:      oldEmbed.URL,
						Channel: targetChannel.ID,
						Embed:   newEmbed,
					})

					if err != nil {
						return fmt.Errorf("error updating embed message in channel %s: %v", targetChannel.Name, err)
					}
				}
			}
		} else {
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
				// 既存のメッセージの中で新しいライブストリームの一覧にないものがあれば、それを削除
				err := s.Session.ChannelMessageDelete(targetChannel.ID, msg.ID)
				if err != nil {
					return fmt.Errorf("error deleting message in channel %s: %v", targetChannel.Name, err)
				}
				// }
			}
		}
	}

	return nil
}

func buildEmbeds(liveStreams []*entities.Video, countryCode string) ([]*discordgo.MessageEmbed, map[string]bool) {
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
