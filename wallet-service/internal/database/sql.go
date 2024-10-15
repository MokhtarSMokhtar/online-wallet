package database

import (
	"database/sql"
	"fmt"
	"github.com/MokhtarSMokhtar/online-wallet/wallet-service/shared"
	_ "github.com/lib/pq"
	"log"
	"sync"
)

var (
	db   *sql.DB
	once sync.Once
)

func GetDB() *sql.DB {
	once.Do(func() {
		cfg := shared.NewConfig()
		conString := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			cfg.DBHost,
			cfg.DBPort,
			cfg.DBUser,
			cfg.DBPassword,
			cfg.DBName,
		)
		var err error
		db, err = sql.Open("postgres", conString)
		if err != nil {
			log.Fatalf("Failed to open database: %v", err)
		}

		// Verify the connection
		if err := db.Ping(); err != nil {
			log.Fatalf("Failed to ping database: %v", err)
		}
	})
	return db
}
func CloseDB() {
	if db != nil {
		if err := db.Close(); err != nil {
			log.Printf("Failed to close database: %v", err)
		}
	}
}
