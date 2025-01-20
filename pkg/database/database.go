package database

import "gorm.io/gorm"

type postgresDB struct {
	db *gorm.DB
}

type PostgresDB interface {
	GetDB() *gorm.DB
	Close()
}
