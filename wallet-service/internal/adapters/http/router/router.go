// internal/adapters/http/router.go

package http

import (
	"github.com/MokhtarSMokhtar/online-wallet/wallet-service/internal/adapters/handlers"
	"github.com/MokhtarSMokhtar/online-wallet/wallet-service/internal/adapters/middlewares"
	"github.com/MokhtarSMokhtar/online-wallet/wallet-service/internal/application/commands"
	"github.com/MokhtarSMokhtar/online-wallet/wallet-service/internal/application/queries"
	"github.com/MokhtarSMokhtar/online-wallet/wallet-service/internal/domain/repositories"
	"github.com/MokhtarSMokhtar/online-wallet/wallet-service/internal/infrastructure/persistence"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Initialize dependencies
	db := persistence.GetDB()
	walletRepo := repositories.NewWalletRepository(db)
	couponRepo := repositories.NewCouponRepository(db)
	commandHandlers := &commands.CommandHandlers{
		TransactionRepo: walletRepo,
		CouponRepo:      couponRepo,
		DB:              db,
	}
	queryHandlers := &queries.QueryHandlers{
		TransactionRepo: walletRepo,
	}

	// Initialize handlers
	commandHandler := handlers.NewCommandHandler(commandHandlers)
	queryHandler := handlers.NewQueryHandler(queryHandlers)

	// Apply middleware
	router.Use(middlewares.AuthMiddleware())

	// Define routes
	walletGroup := router.Group("/wallet")
	{
		walletGroup.GET("/balance", queryHandler.GetBalance)
		walletGroup.GET("/transactions", queryHandler.GetTransactions)
		walletGroup.POST("/transfer", commandHandler.TransferFunds)
		walletGroup.POST("/redeem-coupon", commandHandler.RedeemCoupon)
	}
	return router
}
