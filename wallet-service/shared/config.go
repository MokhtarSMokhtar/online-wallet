package shared

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
	JWTSecret  string
	ISSUER     string
}

func NewConfig() *Config {
	godoterr := godotenv.Load()
	if godoterr != nil {
		log.Fatal("Error loading .env file")
	}
	cfg := &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		JWTSecret:  os.Getenv("JWT_SECRET"),
		ISSUER:     os.Getenv("ISSUER"),
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
