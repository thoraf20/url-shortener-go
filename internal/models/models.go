package models

import "time"

type URL struct {
	TenantID   string            	 `json:"tenant_id"`
	ShortCode  string           	 `json:"short_code"`
	LongUrl    string           		`json:"long_url"`
	CreatedAt  time.Time           `json:"created_at"`
	UpdatedAt  *time.Time          `json:"updated_at"`
}

type ShortenRequest struct {
	TenantID   string            	 `json:"tenant_id"`
	LongUrl    string           		`json:"long_url"`
}

type ShortenResponse struct {
	ShortCode  string           	 `json:"short_code"`
	ShortURL    string           		`json:"long_url"`
}