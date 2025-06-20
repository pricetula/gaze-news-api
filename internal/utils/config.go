package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config holds all app-wide configuration settings
type Config struct {
	DB struct {
		User       string `json:"user"`
		Password   string `json:"password"`
		Name       string `json:"name"`
		Host       string `json:"host"`
		Connection string `json:"connection"`
	} `json:"db"`
	Port         string `json:"port"`
	APP_ENV      string `json:"app_env"`
	NEWS_API_KEY string `json:"news_api_key"`
}

// NewConfig loads configuration from environment variables or a JSON file
func NewConfig() (*Config, error) {
	// Load environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("failed to load .env file: %w", err)
	}

	var cfg Config
	// Unmarshal configuration from environment variables
	err = unmarshal(&cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to parse configuration: %w", err)
	}

	return &cfg, nil
}

func unmarshal(cfg *Config) error {
	// Read from environment variables for simplicity in this example
	// Consider other sources like dedicated config files or libraries in production
	cfg.Port = os.Getenv("PORT")
	cfg.APP_ENV = os.Getenv("APP_ENV")
	cfg.NEWS_API_KEY = os.Getenv("NEWS_API_KEY")
	cfg.DB.User = os.Getenv("DB_USER")
	cfg.DB.Password = os.Getenv("DB_PASSWORD")
	cfg.DB.Name = os.Getenv("DB_NAME")
	cfg.DB.Host = os.Getenv("DB_HOST")
	cfg.DB.Connection = fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=disable",
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Name,
	)

	return nil
}
