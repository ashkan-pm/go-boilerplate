package helpers

import (
	"os"

	"github.com/rs/zerolog/log"

	"github.com/joho/godotenv"
)

func OverloadEnv() {
	env, _ := GetEnv()
	if env == "production" {
		return
	}

	err := godotenv.Load(".env.development")
	if err != nil {
		log.Warn().Err(err).Send()
	}
}

func GetEnv() (env string, httpPort string) {
	env = os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	httpPort = os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	return
}
