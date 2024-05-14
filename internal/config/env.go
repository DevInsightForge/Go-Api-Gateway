package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var (
	once   sync.Once
	config *AppEnvConfig
)

type AppEnvConfig struct {
	ServerAddr string
	ServerPort string
}

func GetAppConfig() *AppEnvConfig {
	return config
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, loading configuration from environment")
	}

	once.Do(loadConfig)
}

func loadConfig() {
	config = &AppEnvConfig{
		ServerAddr: getenv("ADDRESS", "localhost"),
		ServerPort: getenv("PORT", "4000"),
	}
}

func getenv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
