package main

import (
	walletpb "github.com/MokhtarSMokhtar/online-wallet/online-wallet-protos/github.com/MokhtarSMokhtar/online-wallet-protos/wallet"
	_ "github.com/MokhtarSMokhtar/online-wallet/wallet-service/docs"
	grpcserver "github.com/MokhtarSMokhtar/online-wallet/wallet-service/internal/adapters/grpc"
	http "github.com/MokhtarSMokhtar/online-wallet/wallet-service/internal/adapters/http/router"
	"github.com/MokhtarSMokhtar/online-wallet/wallet-service/internal/adapters/messaging"
	"github.com/MokhtarSMokhtar/online-wallet/wallet-service/internal/application/commands"
	"github.com/MokhtarSMokhtar/online-wallet/wallet-service/internal/domain/repositories"
	"github.com/MokhtarSMokhtar/online-wallet/wallet-service/internal/infrastructure"
	"github.com/MokhtarSMokhtar/online-wallet/wallet-service/internal/infrastructure/persistence"
	"github.com/MokhtarSMokhtar/online-wallet/wallet-service/shared"
	"google.golang.org/grpc"
	"log"
	"net"
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
	walletRepo := repositories.NewWalletRepository(db)
	couponRepo := repositories.NewCouponRepository(db)
	commandHandlers := &commands.CommandHandlers{
		TransactionRepo: walletRepo,
		CouponRepo:      couponRepo,
		DB:              db,
	}

	// Start gRPC server
	go func() {
		lis, err := net.Listen("tcp", "localhost:50051")
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}
		grpcServer := grpc.NewServer()
		walletpb.RegisterWalletServiceServer(grpcServer, grpcserver.NewWalletGRPCServer(commandHandlers))
		log.Println("Starting gRPC server on port 50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC server: %v", err)
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
