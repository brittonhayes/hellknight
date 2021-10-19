package basic

import "github.com/bwmarrin/discordgo"

const (
	ID          = "basic-command"
	Description = "Basic command"

	reply = "Hey there! Congratulations, you just executed your first slash command"
)

var ApplicationCommand = &discordgo.ApplicationCommand{
	Name:        ID,
	Description: Description,
}

func Handler() func(*discordgo.Session, *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: reply,
			},
		})
	}
}
