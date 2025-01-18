package main

import (
	"fmt"
	"log"

	"github.com/go-resful-api/configs"
	"github.com/go-resful-api/internal/server"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Initial Config
	appConfig := configs.LoadEnv()

	// Config from pkg
	// email.NewMailProtocol(&appConfig.Mail, branch)
	// pdb := database.SetupPostgresDB(&appConfig.PostgresDB)
	// defer pdb.Close()

	// Initial Server
	app := server.CreateHTTPServer(&appConfig.Server)

	app.Use(cors.New(appConfig.CORS))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	// Run Server
	log.Fatal(app.Listen(fmt.Sprintf(":%d", appConfig.Server.Port)))
}
