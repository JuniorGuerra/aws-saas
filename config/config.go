package config

import (
	"app/internal/models"
	"log"
	"os"
	"strconv"
	"time"
)

type Config struct {
	EmailClient models.EmailClient
	Token       string
}

func LoadConfig() *Config {
	port, err := strconv.Atoi(getEnvOrDefault("SMTP_PORT", "465"))
	if err != nil {
		log.Panic(err)
	}
	timeout, err := strconv.Atoi(getEnvOrDefault("SMTP_TIMEOUT", "5"))
	if err != nil {
		log.Panic(err)
	}
	useSSL, err := strconv.ParseBool(getEnvOrDefault("SMTP_USE_SSL", "true"))
	if err != nil {
		log.Panic(err)
	}

	return &Config{
		EmailClient: models.EmailClient{
			Host:     getEnvOrDefault("SMTP_HOST", "smtp.hostinger.com"),
			Port:     port,
			Username: getEnvOrDefault("SMTP_USERNAME", ""),
			Password: getEnvOrDefault("SMTP_PASSWORD", ""),
			Timeout:  time.Duration(timeout) * time.Second,
			UseSSL:   useSSL,
		},
		Token: getEnvOrDefault("API_TOKEN", ""),
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
