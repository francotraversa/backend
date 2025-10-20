package enviroment

import (
	"github.com/joho/godotenv"
)

type Config struct {
	Port   string
	DB_URL string
}

func LoadEnviroment(key_env string) {
	_ = godotenv.Load(".env." + key_env)
}
