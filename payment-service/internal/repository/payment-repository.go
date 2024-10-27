package repository

import (
	"context"
	"fmt"
	"github.com/MokhatrSMokhtar/online-wallet/payment-service/internal/enums"
	"github.com/MokhatrSMokhtar/online-wallet/payment-service/internal/interfaces"
	"github.com/MokhatrSMokhtar/online-wallet/payment-service/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type paymentRepository struct {
	collection *mongo.Collection
}

func NewPaymentRepository(db *mongo.Database) interfaces.PaymentRepository {
	return &paymentRepository{
		collection: db.Collection("payments"),
	}
}

func (r *paymentRepository) UpdatePaymentRequests(ctx context.Context, payments []*models.PaymentRequest) error {
	var bulkModels []mongo.WriteModel
	for _, payment := range payments {

		filter := bson.M{
			"id": payment.Id,
		}
		update := bson.M{
			"$set": bson.M{
				"payment_status": payment.PaymentStatus,
			},
		}
		updateModel := mongo.NewUpdateOneModel().
			SetFilter(filter).
			SetUpdate(update).
			SetUpsert(false)
		bulkModels = append(bulkModels, updateModel)
	}

	bulkOptions := options.BulkWrite().SetOrdered(false) // Unordered for better performance
	_, err := r.collection.BulkWrite(ctx, bulkModels, bulkOptions)
	if err != nil {
		return fmt.Errorf("failed to bulk update payments: %w", err)
	}

	return nil
}

func (r *paymentRepository) UpdatePaymentRequest(ctx context.Context, payment *models.PaymentRequest) error {
	filter := bson.M{
		"id": payment.Id,
	}
	_, err := r.collection.ReplaceOne(ctx, filter, payment)
	if err != nil {
		return err
	}
	return nil
}

func (r *paymentRepository) GetPaymentRequestByUserAndType(ctx context.Context, userId string, paymentType enums.PaymentType) ([]*models.PaymentRequest, error) {
	var payments []*models.PaymentRequest
	// Define allowed payment statuses
	allowedPaymentStatuses := []enums.PaymentStatus{
		enums.Authorized,
		enums.Initiated,
		enums.Pending,
	}

	// Construct the filter
	filter := bson.M{
		"user_id":        userId,
		"payment_type":   paymentType,
		"payment_status": bson.M{"$in": allowedPaymentStatuses},
	}

	// Execute the Find operation
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to execute Find query: %w", err)
	}
	defer cursor.Close(ctx)

	// Iterate through the cursor
	for cursor.Next(ctx) {
		var payment models.PaymentRequest
		if err := cursor.Decode(&payment); err != nil {
			return nil, fmt.Errorf("failed to decode payment request: %w", err)
		}
		payments = append(payments, &payment)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor encountered error: %w", err)
	}
	return payments, nil
}

func (r *paymentRepository) CreatePayment(ctx context.Context, payment *models.PaymentRequest) error {
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
