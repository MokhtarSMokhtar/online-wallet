package main

import (
	"fmt"
	"github.com/MokhtarSMokhtar/online-wallet/identity-service/config"
	"github.com/MokhtarSMokhtar/online-wallet/identity-service/internal/handlers"
	"log"
	"net/http"
)

func main() {
	port := config.GetPort()

	// Initialize routes
	mux := handlers.InitializeRoutes()

	// Start the server using the mux
	serverAddress := ":" + port
	fmt.Printf("Server is running on port %s\n", port)
	err := http.ListenAndServe(serverAddress, mux)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
