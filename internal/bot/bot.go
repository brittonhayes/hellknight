package bot

import (
	"os"
	"os/signal"

	"github.com/brittonhayes/hellknight/internal/bot/commands/info"
	"github.com/brittonhayes/hellknight/logger"
	"github.com/bwmarrin/discordgo"
)

func Start() {
	log := logger.NewWithPrefix("role", "bot")
	GuildID := os.Getenv("GUILD_ID")
	BotToken := os.Getenv("BOT_TOKEN")

	applicationCommands := []*discordgo.ApplicationCommand{
		info.Command,
	}

	commandHandlers := map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"info": info.Handler,
	}

	s, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		log.Error().Err(err).Msgf("Invalid bot parameters")
	}

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Info().Msg("Bot is up 🚀")
	})

	err = s.Open()
	if err != nil {
		log.Error().Err(err).Msg("Cannot open the session")
	}

	for _, command := range applicationCommands {
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
