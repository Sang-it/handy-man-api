package environment

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	defaultValues = map[string]string{
		"CONNECTION_STRING": "",
		"PORT":              "8080",
	}
)

func Get(key string) string {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	}
	if secret, ok := os.LookupEnv(key); ok {
		return secret
	}
	return defaultValues[key]
}
