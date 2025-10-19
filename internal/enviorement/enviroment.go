package enviroment

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port   string
	DB_URL string
}

func LoadEnviroment(key string) *Config {
	_ = godotenv.Load(".env")
	if env := os.Getenv("ENVIRONMENT"); env != "" {
		_ = godotenv.Overload(".env." + env)
	}
	cfg := &Config{
		Port: os.Getenv("PORT"),
	}

	return cfg
}
