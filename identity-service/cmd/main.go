package main

import (
	"fmt"
	"github.com/MokhtarSMokhtar/online-wallet/identity-service/config"
	_ "github.com/MokhtarSMokhtar/online-wallet/identity-service/docs" // Adjust the import path
	http2 "github.com/MokhtarSMokhtar/online-wallet/identity-service/internal/http/routers"
	"github.com/MokhtarSMokhtar/online-wallet/identity-service/internal/messsage"
	"github.com/MokhtarSMokhtar/online-wallet/identity-service/internal/sql"
	"log"
	"net/http"
)

// @title           Identity Service API
// @version         1.0
// @description     API documentation for the Identity Service
// @termsOfService  http://swagger.io/terms/

// @contact.name    Your Name
// @contact.url     http://www.yourwebsite.com/support

// @license.name    Apache 2.0
// @license.url     http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name
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
