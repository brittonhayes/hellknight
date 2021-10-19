package main

import (
	"os"
	"os/signal"

	"github.com/brittonhayes/hellknight/internal/commands/basic"
	"github.com/brittonhayes/hellknight/internal/logger"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	log = logger.New()
)

func init() {
	godotenv.Load()
}

var (
	commands = []*discordgo.ApplicationCommand{
		basic.ApplicationCommand,
	}

	handlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		basic.ID: basic.Handler(),
	}
)

func main() {
	GuildID := os.Getenv("GUILD_ID")
	BotToken := os.Getenv("BOT_TOKEN")

	s, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		log.Error().Err(err).Msgf("Invalid bot parameters")
	}

	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Info().Msg("Bot is up 🚀")
	})

	err = s.Open()
	if err != nil {
		log.Error().Err(err).Msg("Cannot open the session")
	}

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := handlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	for _, command := range commands {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, GuildID, command)
		if err != nil {
			log.Panic().Msgf("Cannot create '%v' command: %v", command.Name, err)
		}
	}
	defer s.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Info().Msg("Gracefully shutdowning")
}
