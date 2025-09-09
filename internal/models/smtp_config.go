package models

import (
	"time"
)

// SMTPConfig represents a user's SMTP configuration
type SMTPConfig struct {
	ID           string    `json:"id" db:"id"`
	UserID       string    `json:"user_id" db:"user_id"`
	ConfigName   string    `json:"config_name" db:"config_name"`
	APIToken     string    `json:"api_token" db:"api_token"`
	SMTPHost     string    `json:"smtp_host" db:"smtp_host"`
	SMTPPort     int       `json:"smtp_port" db:"smtp_port"`
	SMTPUsername string    `json:"smtp_username" db:"smtp_username"`
	SMTPPassword string    `json:"smtp_password" db:"smtp_password"` // encrypted
	SMTPUseSSL   bool      `json:"smtp_use_ssl" db:"smtp_use_ssl"`
	SMTPTimeout  int       `json:"smtp_timeout" db:"smtp_timeout"`
	IsActive     bool      `json:"is_active" db:"is_active"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

// CreateSMTPConfigRequest represents the request to create a new SMTP config
type CreateSMTPConfigRequest struct {
	ConfigName   string `json:"config_name" validate:"required,min=1,max=255"`
	SMTPHost     string `json:"smtp_host" validate:"required,hostname"`
	SMTPPort     int    `json:"smtp_port" validate:"required,min=1,max=65535"`
	SMTPUsername string `json:"smtp_username" validate:"required,email"`
	SMTPPassword string `json:"smtp_password" validate:"required,min=1"`
	SMTPUseSSL   *bool  `json:"smtp_use_ssl"` // pointer to handle explicit false
	SMTPTimeout  *int   `json:"smtp_timeout" validate:"omitempty,min=1"`
}

// CreateSMTPConfigResponse represents the response after creating SMTP config
type CreateSMTPConfigResponse struct {
	ID         string `json:"id"`
	ConfigName string `json:"config_name"`
	APIToken   string `json:"api_token"`
	CreatedAt  string `json:"created_at"`
}

// ListSMTPConfigsResponse represents the response for listing user's SMTP configs
type ListSMTPConfigsResponse struct {
	Configs []SMTPConfigSummary `json:"configs"`
	Total   int                 `json:"total"`
}

// SMTPConfigSummary represents a summary of SMTP config (without sensitive data)
type SMTPConfigSummary struct {
	ID           string `json:"id"`
	ConfigName   string `json:"config_name"`
	SMTPHost     string `json:"smtp_host"`
	SMTPPort     int    `json:"smtp_port"`
	SMTPUsername string `json:"smtp_username"`
	SMTPUseSSL   bool   `json:"smtp_use_ssl"`
	IsActive     bool   `json:"is_active"`
	CreatedAt    string `json:"created_at"`
}