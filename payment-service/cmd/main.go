package main

import (
	"context"
	"github.com/MokhatrSMokhtar/online-wallet/payment-service/internal/database"
	"github.com/MokhatrSMokhtar/online-wallet/payment-service/internal/enums"
	"github.com/MokhatrSMokhtar/online-wallet/payment-service/internal/models"
	"github.com/MokhatrSMokhtar/online-wallet/payment-service/internal/repository"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	mongoDB := database.GetMongoClient(ctx)
	defer mongoDB.Close(ctx)
	r := repository.NewPaymentRepository(mongoDB.Database)
	payment := models.PaymentRequest{
		UserId:         "user123",
		PaymentType:    enums.Order,
		PaymentMethod:  enums.Credit,
		PaymentStatus:  enums.Initiated,
		Amount:         100.00,
		PaidFromWallet: 50.00,
		IdempotencyKey: "u",
		RequestedAt:    time.Now(),
	}

	err := r.CreatePayment(ctx, payment)
	if err != nil {
		return
	}
}
