package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	Env           string
	Addr          string
	SessionSecret string
	CORSAllowed   []string
	DBDSN         string
	RedisAddr     string
}

func Load() *Config {
	_ = godotenv.Load()
	cfg := &Config{
		Env:           getenv("APP_ENV", "local"),
		Addr:          getenv("APP_ADDR", ":8080"),
		SessionSecret: getenv("SESSION_SECRET", "change-me"),
		DBDSN:         getenv("DB_DSN", ""),
		RedisAddr:     getenv("REDIS_ADDR", "localhost:6379"),
	}
	cors := getenv("CORS_ALLOWED_ORIGINS", "*")
	if cors == "*" {
		cfg.CORSAllowed = []string{"*"}
	} else {
		cfg.CORSAllowed = strings.Split(cors, ",")
	}
	return cfg
}

func getenv(k, d string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return d
}
