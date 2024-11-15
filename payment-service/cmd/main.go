package main

import (
	"github.com/MokhtarSMokhtar/online-wallet/payment-service/docs"
	_ "github.com/MokhtarSMokhtar/online-wallet/payment-service/docs"
	"github.com/MokhtarSMokhtar/online-wallet/payment-service/internal/http/server"
)

func main() {
	docs.SwaggerInfo.Title = "Payment Service API"
	docs.SwaggerInfo.Description = "API documentation for the Payment Service"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/"

	server.InitializePaymentServer()
}
