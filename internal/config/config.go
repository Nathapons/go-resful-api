package config

import (
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

type Config struct {
	Server     ServerConfig
	Mail       MailProtocolConfig
	CORS       cors.Config
	PostgresDB PostgresDBConfig
	Redis      RedisConfig
}

type ServerConfig struct {
	Port               int
	ReadTimeoutSecond  uint
	WriteTimeoutSecond uint
	IdleTimeoutSecond  uint
}

type MailProtocolConfig struct {
	MailerHost      string
	MailerPort      int
	MailerUser      string
	MailerPass      string
	MailerFromEmail string
	MailerToEmail   string
}

type PostgresDBConfig struct {
	PostgresHost string
	PostgresPort string
	PostgresUser string
	PostgresPass string
	PostgresDb   string
}

type RedisConfig struct {
	RedisAddress string
	RedisPort    string
	RedisUser    string
	RedisPass    string
	RedisDb      int
}

func LoadEnv() *Config {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	readTimeout, err := strconv.Atoi(os.Getenv("READ_TIMEOUT_SECOND"))
	if err != nil {
		panic(err)
	}

	writeTimeout, err := strconv.Atoi(os.Getenv("WRITE_TIMEOUT_SECOND"))
	if err != nil {
		panic(err)
	}

	idleTimeout, err := strconv.Atoi(os.Getenv("IDLE_TIMEOUT_SECOND"))
	if err != nil {
		panic(err)
	}

	mailerPort, err := strconv.Atoi(os.Getenv("MAILER_PORT"))
	if err != nil {
		panic(err)
	}

	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DATABASE"))
	if err != nil {
		panic(err)
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8000
	}

	return &Config{
		Server: ServerConfig{
			Port:               port,
			ReadTimeoutSecond:  uint(readTimeout),
			WriteTimeoutSecond: uint(writeTimeout),
			IdleTimeoutSecond:  uint(idleTimeout),
		},
		Mail: MailProtocolConfig{
			MailerHost:      os.Getenv("MAILER_HOST"),
			MailerPort:      mailerPort,
			MailerUser:      os.Getenv("MAILER_USER"),
			MailerPass:      os.Getenv("MAILER_PASS"),
			MailerFromEmail: os.Getenv("MAILER_FROM_EMAIL"),
			MailerToEmail:   os.Getenv("MAILER_TO_EMAIL"),
		},
		CORS: cors.Config{
			AllowOrigins: os.Getenv("CORS_ALLOW_ORIGINS"),
			AllowMethods: os.Getenv("CORS_ALLOW_METHODS"),
			AllowHeaders: os.Getenv("CORS_ALLOW_HEADERS"),
		},
		PostgresDB: PostgresDBConfig{
			PostgresHost: os.Getenv("POSTGRES_HOST"),
			PostgresPort: os.Getenv("POSTGRES_PORT"),
			PostgresUser: os.Getenv("POSTGRES_USER"),
			PostgresPass: os.Getenv("POSTGRES_PASS"),
			PostgresDb:   os.Getenv("POSTGRES_DB"),
		},
		Redis: RedisConfig{
			RedisAddress: os.Getenv("REDIS_ADDRESS"),
			RedisPort:    os.Getenv("REDIS_PORT"),
			RedisUser:    os.Getenv("REDIS_USERNAME"),
			RedisPass:    os.Getenv("REDIS_PASSWORD"),
			RedisDb:      redisDB,
		},
	}
}
