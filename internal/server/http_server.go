package server

import (
	"time"

	"github.com/go-resful-api/configs"
	"github.com/gofiber/fiber/v2"
)

func CreateHTTPServer(serverConfig *configs.ServerConfig) *fiber.App {
	app := fiber.New(fiber.Config{
		ReadTimeout:  time.Duration(serverConfig.ReadTimeoutSecond) * time.Second,
		WriteTimeout: time.Duration(serverConfig.WriteTimeoutSecond) * time.Second,
		IdleTimeout:  time.Duration(serverConfig.IdleTimeoutSecond) * time.Second,
	})
	return app
}
