package server

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/helmet/v2"
	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load()
}

var config = fiber.Config{
	AppName:               "hellknight-server",
	GETOnly:               true,
	DisableStartupMessage: true,
}

func setupMiddlewares(app *fiber.App) {
	app.Use(helmet.New())
	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))
	app.Use(etag.New())
	if os.Getenv("ENABLE_LIMITER") != "" {
		app.Use(limiter.New())
	}
	if os.Getenv("ENABLE_LOGGER") != "" {
		app.Use(logger.New())
	}
}

func Create() *fiber.App {
	app := fiber.New(config)
	setupMiddlewares(app)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	return app
}

func Listen(app *fiber.App) error {
	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	serverPort := os.Getenv("PORT")
	return app.Listen(fmt.Sprintf(":%s", serverPort))
}
