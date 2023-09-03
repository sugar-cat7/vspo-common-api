package main

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

func main() {
	token := os.Getenv("DISCORD_BOT_TOKEN")

	// Discordセッションの作成
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	botUser, err := dg.User("@me")
	if err != nil {
		fmt.Println("error getting bot user,", err)
		return
	}

	// 所属する全てのサーバー（ギルド）の取得
	guilds, err := dg.UserGuilds(100, "", "")
	if err != nil {
		fmt.Println("error getting user guilds,", err)
		return
	}

	// チャンネルの名前を定義
	targetChannelName := "general"

	// 初回メッセージの内容
	initialMessage := "すぽじゅーるは、ぶいすぽっ!メンバーの配信(Youtube/Twitch/ツイキャス/ニコニコ)や切り抜きを一覧で確認できる非公式サイトです。 /Spodule aggregates schedules for Japan's Vtuber group, Vspo.\n\nWeb版はこちら：https://www.vspo-schedule.com/schedule/all"

	// 各サーバーの特定のチャンネルにメッセージを送信
	for _, guild := range guilds {
		channels, err := dg.GuildChannels(guild.ID)
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
			targetChannel, err = dg.GuildChannelCreate(guild.ID, targetChannelName, discordgo.ChannelTypeGuildText)
			if err != nil {
				fmt.Printf("error creating channel %s: %v\n", targetChannelName, err)
				continue
			}
			_, err = dg.ChannelMessageSend(targetChannel.ID, initialMessage)
			if err != nil {
				fmt.Printf("error sending initial message to channel %s: %v\n", targetChannel.Name, err)
			}
		} else {
			messages, err := dg.ChannelMessages(targetChannel.ID, 100, "", "", "")
			if err != nil {
				fmt.Printf("error getting messages from channel %s: %v\n", targetChannel.Name, err)
				continue
			}

			var isFirstMessage = true
			for _, message := range messages {
				if message.Author.ID == botUser.ID {
					if isFirstMessage {
						isFirstMessage = false
						continue
					}
					err = dg.ChannelMessageDelete(targetChannel.ID, message.ID)
					if err != nil {
						fmt.Printf("error deleting message in channel %s: %v\n", targetChannel.Name, err)
					}
				}
			}
			_, err = dg.ChannelMessageSend(targetChannel.ID, "新しい挨拶メッセージ")
			if err != nil {
				fmt.Printf("error sending message to channel %s: %v\n", targetChannel.Name, err)
			}
		}
	}
}
