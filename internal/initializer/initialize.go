package initializer

import (
	"log"

	"github.com/joho/godotenv"
)

func InitializeEnvironmentVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error occurred while initializing environment variables.")
	}
}
