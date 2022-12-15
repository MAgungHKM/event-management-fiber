package env

import (
	"event-management/utils"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Setup() {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func Get(key string, defaultValue ...string) string {
	value := os.Getenv(key)

	if value == "" && len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return value
}

func GetInt(key string, def ...int) int {
	value := -1
	valuePoint := utils.ParseInt(os.Getenv(key))
	if valuePoint != nil {
		value = *valuePoint
	}

	if value == -1 && len(def) > 0 {
		value = def[0]
	}

	return value
}
