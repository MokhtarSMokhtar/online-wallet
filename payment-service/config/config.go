package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DBHost       string `json:"db_host"`
	DBPort       string `json:"db_port"`
	DBUser       string `json:"db_user"`
	DBPassword   string `json:"db_password"`
	DBName       string `json:"db_name"`
	TABPublicKey string `json:"tab_public_key"`
	TABSecretKey string `json:"tab_secret_key"`
	PostUrl      string `json:"post_url"`
	BaseURl      string `json:"base_url"`
}

func NewConfig() *Config {
	godoterr := godotenv.Load()
	if godoterr != nil {
		log.Fatal("Error loading .env file")
	}
	cfg := &Config{
		DBHost:       os.Getenv("DB_HOST"),
		DBPort:       os.Getenv("DB_PORT"),
		DBUser:       os.Getenv("MONGO_INITDB_ROOT_USERNAME"),
		DBPassword:   os.Getenv("MONGO_INITDB_ROOT_PASSWORD"),
		DBName:       os.Getenv("DB_NAME"),
		TABPublicKey: os.Getenv("TAB_PUBLIC_KEY"),
		TABSecretKey: os.Getenv("TAB_SECRET_KEY"),
		PostUrl:      os.Getenv("POST_URL"),
		BaseURl:      os.Getenv("LOCAL_URL"),
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
