package persistence

import (
	"fmt"
	"github.com/MokhtarSMokhtar/online-wallet/wallet-service/shared"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"sync"
)

var (
	db   *gorm.DB
	once sync.Once
)

func GetDB() *gorm.DB {
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
		db, err = gorm.Open(postgres.Open(conString), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to open database: %v", err)
		}
	})
	return db
}
func CloseDB() error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
