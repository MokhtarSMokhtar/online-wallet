package server

import (
	"context"
	"github.com/MokhtarSMokhtar/online-wallet/comman/middelwares"
	"github.com/MokhtarSMokhtar/online-wallet/payment-service/config"
	"github.com/MokhtarSMokhtar/online-wallet/payment-service/internal/database"
	"github.com/MokhtarSMokhtar/online-wallet/payment-service/internal/handler"
	"github.com/MokhtarSMokhtar/online-wallet/payment-service/internal/repository"
	tapclient "github.com/MokhtarSMokhtar/online-wallet/payment-service/internal/tap-payment/http"
	services "github.com/MokhtarSMokhtar/online-wallet/payment-service/internal/tap-payment/tapservice"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
)

func InitializePaymentServer() {
	// Load configuration
	cfg := config.NewConfig()

	// Initialize database connection
	ctx := context.Background()

	db := database.GetMongoClient(ctx)
	defer db.Close(ctx)

	// Initialize repositories
	paymentRepo := repository.NewPaymentRepository(db.Database)

	tapClient := tapclient.NewHttpClientFactory(cfg.TABSecretKey, "v2")

	paymentService := services.NewPaymentService(tapClient)

	paymentHandler := handler.NewPaymentHandler(paymentService, paymentRepo)

	mux := http.NewServeMux()
	mux.HandleFunc("/swagger/", httpSwagger.WrapHandler)
	mux.Handle("/payments", middelwares.AuthMiddleware(http.HandlerFunc(paymentHandler.UserPaymentHandler)))
	mux.Handle("/payments/capture", http.HandlerFunc(paymentHandler.CapturePayment))

	port := config.GetPort()
	// Start the server
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Starting server on port %s...", port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on %s: %v\n", port, err)
	}

}
