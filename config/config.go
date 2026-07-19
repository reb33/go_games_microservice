package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DSN string
}

func NewConfig() *Config {
	// Игнорируем ошибку. Если .env файла нет, это нормально (например, в Docker)
	// Приложение просто прочитает системные переменные окружения.
	_ = godotenv.Load(".env")

	dsn := os.Getenv("DSN")
	if dsn == "" {
		panic("CRITICAL: DSN environment variable is not set")
	}
	return &Config{
		DSN: dsn,
	}
}
