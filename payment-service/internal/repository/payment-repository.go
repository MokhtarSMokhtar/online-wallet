package repository

import (
	"context"
	"github.com/MokhatrSMokhtar/online-wallet/payment-service/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PaymentRepository interface {
	CreatePayment(ctx context.Context, payment models.PaymentRequest) error
	GetPaymentByIdempotencyKey(ctx context.Context, key string) (*models.PaymentRequest, error)
}
type paymentRepository struct {
	collection *mongo.Collection
}

func NewPaymentRepository(db *mongo.Database) PaymentRepository {
	return &paymentRepository{
		collection: db.Collection("payments"),
	}
}

func (r *paymentRepository) CreatePayment(ctx context.Context, payment models.PaymentRequest) error {
	_, err := r.collection.InsertOne(ctx, payment)
	return err
}

func (r *paymentRepository) GetPaymentByIdempotencyKey(ctx context.Context, key string) (*models.PaymentRequest, error) {
	var payment models.PaymentRequest
	err := r.collection.FindOne(ctx, bson.M{"idempotency_key": key}).Decode(&payment)
	if err != nil {
		return nil, err
	}
	return &payment, nil
}
