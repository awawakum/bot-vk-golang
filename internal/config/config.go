package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	VkConfig
	DBConfig
}

type VkConfig struct {
	VkToken string
}

type DBConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

func NewConfig() (*Config, error) {

	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	tg := VkConfig{
		VkToken: os.Getenv("TOKEN"),
	}

	db := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
	}

	return &Config{
		DBConfig: db,
		VkConfig: tg,
	}, nil
}
