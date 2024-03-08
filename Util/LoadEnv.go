package Util

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return err
	}
	return nil
}

func GetEnvVariable(key string) string {
	return os.Getenv(key)
}
