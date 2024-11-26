package config

import (
	"log"
	"os"
)

var JWTSecret string

// LoadConfig は環境変数を読み込む
func LoadConfig() {
	JWTSecret = os.Getenv("JWT_SECRET_KEY")
	if JWTSecret == "" {
		log.Fatal("JWT_SECRET_KEY is not set")
	}
}
