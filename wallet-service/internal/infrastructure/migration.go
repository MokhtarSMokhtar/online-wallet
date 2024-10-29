package infrastructure

import (
	"github.com/MokhtarSMokhtar/online-wallet/wallet-service/internal/domain/models"
	"github.com/MokhtarSMokhtar/online-wallet/wallet-service/internal/infrastructure/persistence"
	"log"
)

func MigrateDb() {
	db := persistence.GetDB()
	err := db.AutoMigrate(&models.WalletTransaction{}, &models.Coupon{}, &models.CouponUsage{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database migrated successfully")
}
