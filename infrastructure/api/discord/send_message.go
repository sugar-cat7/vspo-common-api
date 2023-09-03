package ports

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/sugar-cat7/vspo-common-api/domain/entities"
	"github.com/sugar-cat7/vspo-common-api/domain/ports"
	"github.com/sugar-cat7/vspo-common-api/util"
)

type discordServiceImpl struct {
	Session *discordgo.Session
}

// NewDiscordService creates a new DiscordService.
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
		fmt.Println("error getting user guilds,", err)
		return fmt.Errorf("error getting user guilds: %v", err)
	}

	// チャンネルの名前を定義
	targetChannelName := "ぶいすぽ配信情報"

	// 初回メッセージの内容
	initialMessage := "すぽじゅーるは、ぶいすぽっ!メンバーの配信(Youtube/Twitch/ツイキャス/ニコニコ)や切り抜きを一覧で確認できる非公式サイトです。 /Spodule aggregates schedules for Japan's Vtuber group, Vspo.\n\nWeb版はこちら：https://www.vspo-schedule.com/schedule/all"

	// 各サーバーの特定のチャンネルにメッセージを送信
	for _, guild := range guilds {
		channels, err := s.Session.GuildChannels(guild.ID)
		if err != nil {
			fmt.Println("error getting guild channels,", err)
			continue
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
				fmt.Printf("error creating channel %s: %v\n", targetChannelName, err)
				continue
			}
			_, err = s.Session.ChannelMessageSend(targetChannel.ID, initialMessage)
			if err != nil {
				fmt.Printf("error sending initial message to channel %s: %v\n", targetChannel.Name, err)
			}
		}

		var embeds []*discordgo.MessageEmbed
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
		}
		messages, err := s.Session.ChannelMessages(targetChannel.ID, 100, "", "", "")
		if err != nil {
			fmt.Printf("error getting messages from channel %s: %v\n", targetChannel.Name, err)
			continue
		}

		for _, message := range messages {
			if message.Author.ID == botUser.ID && message.Content != initialMessage {
				err = s.Session.ChannelMessageDelete(targetChannel.ID, message.ID)
				if err != nil {
					fmt.Printf("error deleting message in channel %s: %v\n", targetChannel.Name, err)
				}
			}
		}

		var currentEmbeds []*discordgo.MessageEmbed
		currentSize := 0
		const maxEmbedSize = 6000
		for _, embed := range embeds {
			embedSize := util.CalculateEmbedSize(embed)
			if currentSize+embedSize > maxEmbedSize || len(currentEmbeds) == 10 {
				_, err = s.Session.ChannelMessageSendComplex(targetChannel.ID, &discordgo.MessageSend{
					Embeds: currentEmbeds,
				})
				if err != nil {
					fmt.Printf("error sending embed message to channel %s: %v\n", targetChannel.Name, err)
				}
				currentEmbeds = []*discordgo.MessageEmbed{}
				currentSize = 0
			}
			currentEmbeds = append(currentEmbeds, embed)
			currentSize += embedSize
		}

		if len(currentEmbeds) > 0 {
			_, err = s.Session.ChannelMessageSendComplex(targetChannel.ID, &discordgo.MessageSend{
				Embeds: currentEmbeds,
			})
			if err != nil {
				fmt.Printf("error sending embed message to channel %s: %v\n", targetChannel.Name, err)
			}
		}

	}
	return nil
}
