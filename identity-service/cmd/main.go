package main

import (
	"fmt"
	"github.com/MokhtarSMokhtar/online-wallet/identity-service/config"
	http2 "github.com/MokhtarSMokhtar/online-wallet/identity-service/internal/http/routers"
	"github.com/MokhtarSMokhtar/online-wallet/identity-service/internal/messsage"
	"github.com/MokhtarSMokhtar/online-wallet/identity-service/internal/sql"
	"log"
	"net/http"
)

func main() {
	port := config.GetPort()

	// Initialize routes
	mux := http2.InitializeRoutes()
	rabbitMQ := messsage.GetRabbitMQInstance()
	defer rabbitMQ.Close()
	s := sql.MigrateDatabase()
	if s != nil {
		return
	}
	// Start the server using the mux
	serverAddress := ":" + port
	fmt.Printf("Server is running on port %s\n", port)
	err := http.ListenAndServe(serverAddress, mux)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
