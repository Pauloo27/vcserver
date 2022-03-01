package bootstrap

import (
	"github.com/Pauloo27/logger"
	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		logger.Fatal(err)
	}
}
