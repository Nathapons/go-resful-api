package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-resful-api/internal/configs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Initial Config
	cfg := configs.LoadEnv()

	// Initial Server
	app := fiber.New(fiber.Config{
		ReadTimeout:  time.Duration(cfg.Server.ReadTimeoutSecond) * time.Second,
		WriteTimeout: time.Duration(cfg.Server.WriteTimeoutSecond) * time.Second,
		IdleTimeout:  time.Duration(cfg.Server.IdleTimeoutSecond) * time.Second,
	})

	app.Use(cors.New(cfg.CORS))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	// Run Server
	log.Fatal(app.Listen(fmt.Sprintf(":%d", cfg.Server.Port)))
}
