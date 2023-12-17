package config

import (
	"os"
	"time"
)

// AppConfig holds the configuration for the application
type AppConfig struct {
	FilePath      string
	APIKey        string
	Email         string
	Domain        string
	Subdomain     string
	CheckInterval time.Duration
}

// LoadConfig loads configuration settings from environment variables
func LoadConfig() (*AppConfig, error) {
	checkInterval, err := time.ParseDuration(os.Getenv("CHECK_INTERVAL"))
	if err != nil {
		return nil, err
	}

	return &AppConfig{
		FilePath:      os.Getenv("FILE_PATH"),
		APIKey:        os.Getenv("API_KEY"),
		Email:         os.Getenv("EMAIL"),
		Domain:        os.Getenv("DOMAIN"),
		Subdomain:     os.Getenv("SUBDOMAIN"),
		CheckInterval: checkInterval,
	}, nil
}
