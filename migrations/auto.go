package main

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	postgresDB "go-games-microservice/internal/repository/postgres"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&postgresDB.MatchDB{}, &postgresDB.PlayerDB{}, &postgresDB.MatchEventDB{})
}
