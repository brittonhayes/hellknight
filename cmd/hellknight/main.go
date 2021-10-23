package main

import (
	"github.com/brittonhayes/hellknight/internal/bot"
	"github.com/brittonhayes/hellknight/internal/server"
	"github.com/brittonhayes/hellknight/logger"
	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load()
}

func main() {
	go func() {
		app := server.Create()
		log := logger.NewWithPrefix("role", "api")
		log.Info().Msg("API server is up 🚀")
		if err := server.Listen(app); err != nil {
			log.Panic().Err(err)
		}
	}()
	bot.Start()
}
