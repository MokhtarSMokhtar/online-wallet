package handlers

import (
	"github.com/MokhtarSMokhtar/online-wallet/identity-service/internal/repository"
	"github.com/MokhtarSMokhtar/online-wallet/identity-service/internal/sql"
	"log"
	"net/http"
)

func InitializeRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// Initialize database connection
	db := sql.NewIdentity()
	conn, err := db.GetConnection()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize repository
	userRepo := &repository.UserRepository{
		Db: conn,
	}

	// Initialize handler
	userHandler := NewUserHandler(userRepo)

	// Define routes and associate them with handlers
	mux.HandleFunc("/signup", userHandler.Signup)
	mux.HandleFunc("/login", userHandler.Login)
	return mux
}
