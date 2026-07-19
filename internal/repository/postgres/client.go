package postgres

import (
	"go-games-microservice/config"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	*gorm.DB
}

func NewDB(conf config.Config) (*DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)
	db, err := gorm.Open(postgres.Open(conf.DSN), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err.Error())
	}
	return &DB{db}, nil
}
