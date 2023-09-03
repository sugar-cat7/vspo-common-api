package util

import "github.com/bwmarrin/discordgo"

func CalculateEmbedSize(embed *discordgo.MessageEmbed) int {
	size := len(embed.Title) + len(embed.Description) + len(embed.URL)
	for _, field := range embed.Fields {
		size += len(field.Name) + len(field.Value)
	}
	if embed.Author != nil {
		size += len(embed.Author.Name)
	}
	if embed.Image != nil {
		size += len(embed.Image.URL)
	}
	return size
}
