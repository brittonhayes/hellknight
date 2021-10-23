package info

import (
	"github.com/brittonhayes/hellknight/logger"
	"github.com/bwmarrin/discordgo"
)

const (
	NAME = "info"
	DESC = "Return application information"

	reply = "Thanks for using Hellknight!👋 \n \nIf this project helped you, please consider starring it on github:\n ⭐ https://github.com/brittonhayes/hellknight"
)

var Command = &discordgo.ApplicationCommand{
	ID:          NAME,
	Name:        NAME,
	Description: DESC,
}

func Handler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	log := logger.NewWithPrefix("cmd", NAME)
	log.Info().Msgf("executing %q command", NAME)
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: reply,
		},
	})

}
