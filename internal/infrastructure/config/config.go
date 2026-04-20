package config

import "os"

type Config struct {
	AccessToken   string
	BusinessID    string
	PhoneNumberID string
	BaseURL       string
}

// Load returns a Config object with values set from environment variables.
func Load() *Config {
	return &Config{
		AccessToken:   os.Getenv("META_ACCESS_TOKEN"),
		BusinessID:    os.Getenv("META_BUSINESS_ID"),
		PhoneNumberID: os.Getenv("META_PHONE_NUMBER_ID"),
		BaseURL:       "https://graph.facebook.com/v18.0",
	}
}