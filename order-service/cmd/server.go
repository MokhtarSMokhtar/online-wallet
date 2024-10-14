package main

import (
	"github.com/MokhtarSMokhtar/online-wallet/order-service/config"
	"github.com/MokhtarSMokhtar/online-wallet/order-service/internal/handler"
	"github.com/MokhtarSMokhtar/online-wallet/order-service/internal/middelwares"
	"net/http"
)

func initializeServer(orderHandler *handler.OrderHandler) *http.Server {
	mux := http.NewServeMux()

	// Routes
	mux.HandleFunc("/orders", orderHandler.CreateOrder)
	mux.Handle("/orders", middelwares.AuthMiddleware(mux))
	return &http.Server{
		Addr:    ":" + config.GetPort(),
		Handler: mux,
	}
}
