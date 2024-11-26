package config

import (
	"log"
	"os"
)

// Config holds all configuration values
type Config struct {
	JWTSecret string
}

// Global instance of Config
var cfg *Config

// GetConfig returns the current configuration
func GetConfig() *Config {
	if cfg == nil {
		panic("Configuration not initialized. Ensure LoadConfig() is called first")
	}
	return cfg
}

// LoadConfig は環境変数を読み込む
func LoadConfig() {
	const minSecretLength = 32
	secret := os.Getenv("JWT_SECRET_KEY")
	if secret == "" {
		log.Printf("ERROR: JWT_SECRET_KEY is not set")
		panic("JWT_SECRET_KEY is required for application security")
	}

	if len(secret) < minSecretLength {
		log.Printf("ERROR: JWT_SECRET_KEY must be at least %d characters long", minSecretLength)
		panic("JWT_SECRET_KEY is too short")

	}

	cfg = &Config{JWTSecret: secret}
	log.Printf("Configuration loaded successfully")
}
