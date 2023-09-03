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

	for _, guild := range guilds {
		if err := s.processGuild(guild, liveStreams, botUser, countryCode); err != nil {
			fmt.Printf("error processing guild %s: %v\n", guild.Name, err)
			continue
		}
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
		_, err = s.Session.ChannelMessageSend(targetChannel.ID, initialMessage)
		if err != nil {
			return fmt.Errorf("error sending initial message to channel %s: %v", targetChannel.Name, err)
		}
	}

	embeds, isExistVideoMap := buildEmbeds(liveStreams, countryCode)
	messages, err := s.Session.ChannelMessages(targetChannel.ID, 100, "", "", "")
	if err != nil {
		return fmt.Errorf("error getting messages from channel %s: %v", targetChannel.Name, err)
	}

	var isExistChannelVideoMap = make(map[string]bool)
	// Process the messages for the given channel.
	for _, message := range messages {
		if message.Author.ID == botUser.ID && message.Content != initialMessage {
			for _, embed := range message.Embeds {
				isExistChannelVideoMap[embed.URL] = true
				if embed.URL != "" && !isExistVideoMap[embed.URL] {
					err = s.Session.ChannelMessageDelete(targetChannel.ID, message.ID)
					if err != nil {
						return fmt.Errorf("error deleting message in channel %s: %v", targetChannel.Name, err)
					}
				} else {
					for _, newEmbed := range embeds {
						if embed.URL == newEmbed.URL {
							if embed.Title != newEmbed.Title || embed.Fields[0].Value != newEmbed.Fields[0].Value {
								_, err := s.Session.ChannelMessageEditComplex(&discordgo.MessageEdit{
									ID:      message.ID,
									Channel: targetChannel.ID,
									Embed:   newEmbed,
								})
								if err != nil {
									return fmt.Errorf("error updating message in channel %s: %v", targetChannel.Name, err)
								}
							}
							break
						}
					}
				}
			}
		}
	}

	var currentEmbeds []*discordgo.MessageEmbed
	currentSize := 0
	const maxEmbedSize = 6000

	// Split the embeds to ensure they fit within Discord's message limits.
	for _, embed := range embeds {
		if !isExistChannelVideoMap[embed.URL] {
			embedSize := util.CalculateEmbedSize(embed)
			if currentSize+embedSize > maxEmbedSize || len(currentEmbeds) == 10 {
				_, err = s.Session.ChannelMessageSendComplex(targetChannel.ID, &discordgo.MessageSend{
					Embeds: currentEmbeds,
				})
				if err != nil {
					return fmt.Errorf("error sending embed message to channel %s: %v", targetChannel.Name, err)
				}
				currentEmbeds = []*discordgo.MessageEmbed{}
				currentSize = 0
			}
			currentEmbeds = append(currentEmbeds, embed)
			currentSize += embedSize
		}
	}

	// Ensure any remaining embeds are also sent.
	if len(currentEmbeds) > 0 {
		_, err = s.Session.ChannelMessageSendComplex(targetChannel.ID, &discordgo.MessageSend{
			Embeds: currentEmbeds,
		})
		if err != nil {
			return fmt.Errorf("error sending embed message to channel %s: %v", targetChannel.Name, err)
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
