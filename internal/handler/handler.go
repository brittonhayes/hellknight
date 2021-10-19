package handler

import "github.com/bwmarrin/discordgo"

func Add(name string, handler func(s *discordgo.Session, i *discordgo.InteractionCreate), dest map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)) {
	dest[name] = handler
}
