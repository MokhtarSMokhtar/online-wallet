package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	JWTSecret string
	ISSUER    string
}

func NewConfig() *Config {
	godoterr := godotenv.Load()
	if godoterr != nil {
		log.Fatal("Error loading .env file")
	}
	cfg := &Config{
		JWTSecret: os.Getenv("JWT_SECRET"),
		ISSUER:    os.Getenv("ISSUER"),
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
