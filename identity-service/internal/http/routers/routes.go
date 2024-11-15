package routers

import (
	"github.com/MokhtarSMokhtar/online-wallet/identity-service/internal/handlers"
	"github.com/MokhtarSMokhtar/online-wallet/identity-service/internal/http/middelwares"
	"github.com/MokhtarSMokhtar/online-wallet/identity-service/internal/repository"
	"github.com/MokhtarSMokhtar/online-wallet/identity-service/internal/sql"
	httpSwagger "github.com/swaggo/http-swagger"
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
	userHandler := handlers.NewUserHandler(userRepo)

	// Define routes and associate them with handlers
	mux.HandleFunc("/signup", userHandler.Signup)
	mux.HandleFunc("/login", userHandler.Login)
	mux.HandleFunc("/swagger/", httpSwagger.WrapHandler)
	protectedMux := http.NewServeMux()
	protectedMux.HandleFunc("/protected", userHandler.ProtectedEndpoint)

	// Apply middleware to protected routes
	mux.Handle("/protected", middelwares.AuthMiddleware(protectedMux))
	return mux
}
