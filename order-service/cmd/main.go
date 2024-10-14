package main

import (
	"context"
	"github.com/MokhtarSMokhtar/online-wallet/order-service/config"
	"github.com/MokhtarSMokhtar/online-wallet/order-service/internal/database/mongo"
	"github.com/MokhtarSMokhtar/online-wallet/order-service/internal/handler"
	"github.com/MokhtarSMokhtar/online-wallet/order-service/internal/repository"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg := config.NewConfig()
	mongodb := mongo.OrderDB{}
	port := config.GetPort()
	err := mongodb.ConnectToDb()
	if err != nil {
		return
	}
	defer func() {
		if err := mongodb.Client.Disconnect(context.Background()); err != nil {
			log.Fatalf("Error disconnecting from MongoDB: %v", err)
		}
	}()

	// Initialize Repository
	orderRepo := repository.NewOrderRepository(mongodb.Client, cfg.DBName)

	// Initialize Handlers
	orderHandler := handler.NewOrderHandler(orderRepo)

	// Initialize Server and Routes
	ser := initializeServer(orderHandler)

	// Start Server
	go func() {
		log.Printf("Server is running on port %s", port)
		if err := ser.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on port %s: %v", port, err)
		}
	}()

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := ser.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")

}
