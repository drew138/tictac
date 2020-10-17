package environment

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var envVariables = []string{
	"DB_USERNAME",
	"DB_PASSWORD",
	"DB_NAME",
	"DB_HOST",
	"DB_PORT",
	"JWT_SECRET_KEY",
	"JWT_REFRESH_SECRET_KEY",
}

func checkVars() bool {
	var exists bool
	for i := 0; i < len(envVariables); i++ {
		_, exists = os.LookupEnv(envVariables[i])
		if !exists {
			return exists
		}
	}
	return exists
}

// LoadEnv - check if env vars are already in os, if not load .env
func LoadEnv() {
	var err error
	if !checkVars() {
		err = godotenv.Load(".env")
	}
	if err != nil || !checkVars() {
		log.Fatal("Error loading .env file")
	}
}
