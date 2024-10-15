package main

import (
	"github.com/MokhtarSMokhtar/online-wallet/wallet-service/internal/database"
	"github.com/MokhtarSMokhtar/online-wallet/wallet-service/internal/messaging"
	"github.com/MokhtarSMokhtar/online-wallet/wallet-service/shared"
	"log"
	"net/http"
)

func main() {
	defer database.CloseDB()
	rabbitMQ := messaging.GetRabbitMQInstance()
	defer rabbitMQ.Close()

	// Run database migrations

	// Start consuming UserRegistered events
	rabbitMQ.ConsumeUserRegisteredEvents()

	// Start the server
	port := shared.GetPort()
	log.Printf("Wallet Service is running on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
