package main

import (
	_ "github.com/MokhtarSMokhtar/online-wallet/wallet-service/docs" // Import the generated docs
	http "github.com/MokhtarSMokhtar/online-wallet/wallet-service/internal/adapters/http/router"
	"github.com/MokhtarSMokhtar/online-wallet/wallet-service/internal/adapters/messaging"
	"github.com/MokhtarSMokhtar/online-wallet/wallet-service/internal/infrastructure"
	"github.com/MokhtarSMokhtar/online-wallet/wallet-service/internal/infrastructure/persistence"
	"github.com/MokhtarSMokhtar/online-wallet/wallet-service/shared"
	"log"
)

// @title           Online Wallet API
// @version         1.0
// @description     API documentation for the Online Wallet Service

func main() {
	db := persistence.GetDB()

	defer func() {
		sqlDB, err := db.DB()
		if err == nil {
			sqlDB.Close()
		}
	}()

	// Run migrations
	infrastructure.MigrateDb()

	rabbitMQ := messaging.GetRabbitMQInstance()
	defer rabbitMQ.Close()
	infrastructure.MigrateDb()
	rabbitMQ.ConsumeUserRegisteredEvents()
	port := shared.GetPort()
	// Start the server
	router := http.SetupRouter()
	log.Printf("Starting server on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
