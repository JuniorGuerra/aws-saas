package config

import (
	"log"
	"os"
	"strconv"
	"time"
)

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

type EmailClient struct {
	Host     string
	Port     int
	Username string
	Password string
	Timeout  time.Duration // time.Duration
	UseSSL   bool
}

type Config struct {
	EmailClient EmailClient
	Token       string
	Database    PostgresConfig
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
		EmailClient: EmailClient{
			Host:     getEnvOrDefault("SMTP_HOST", "smtp.hostinger.com"),
			Port:     port,
			Username: getEnvOrDefault("SMTP_USERNAME", ""),
			Password: getEnvOrDefault("SMTP_PASSWORD", ""),
			Timeout:  time.Duration(timeout) * time.Second,
			UseSSL:   useSSL,
		},
		Token: getEnvOrDefault("API_TOKEN", ""),
		Database: PostgresConfig{
			Host:     getEnvOrDefault("DB_HOST", "localhost"),
			Port:     getEnvOrDefault("DB_PORT", "5432"),
			User:     getEnvOrDefault("DB_USER", "postgres"),
			Password: getEnvOrDefault("DB_PASSWORD", ""),
			Database: getEnvOrDefault("DB_NAME", "postgres"),
			SSLMode:  getEnvOrDefault("DB_SSLMODE", "require"),
		},
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
