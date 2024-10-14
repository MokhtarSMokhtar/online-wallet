package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func NewConfig() *Config {
	godoterr := godotenv.Load()
	if godoterr != nil {
		log.Fatal("Error loading .env file")
	}
	cfg := &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("MONGO_INITDB_ROOT_USERNAME"),
		DBPassword: os.Getenv("MONGO_INITDB_ROOT_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
	}
	return cfg
}
func GetPort() string {

	godoterr := godotenv.Load()
	if godoterr != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	return port
}
