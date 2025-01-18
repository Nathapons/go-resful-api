package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-resful-api/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type postgresDB struct {
	db *gorm.DB
}

type PostgresDB interface {
	GetDB() *gorm.DB
	Close()
}

func SetupPostgresDB(configs *configs.PostgresDBConfig) PostgresDB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		configs.PostgresHost,
		configs.PostgresPort,
		configs.PostgresUser,
		configs.PostgresPass,
		configs.PostgresDb,
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		log.Println("Error occurred while connecting to the database:", err)
		log.Fatalln("Failed to connect to Postgres Database!!")
	}

	fmt.Println("Connect Postgres Database Complete!!")
	return &postgresDB{db: db}
}

func (pdb *postgresDB) GetDB() *gorm.DB {
	return pdb.db
}

func (pdb *postgresDB) Close() {
	sqlDB, err := pdb.db.DB()
	if err != nil {
		log.Println("Error getting database handle:", err)
		return
	}
	if err := sqlDB.Close(); err != nil {
		log.Println("Error closing database connection:", err)
	}
	fmt.Println("Postgres Database connection closed.")
}
