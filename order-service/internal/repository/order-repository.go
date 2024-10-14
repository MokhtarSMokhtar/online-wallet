package repository

import (
	"context"
	"github.com/MokhtarSMokhtar/online-wallet/order-service/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository struct {
	Collection *mongo.Collection
}

func NewOrderRepository(client *mongo.Client, dbName string) *OrderRepository {
	collection := client.Database(dbName).Collection("orders")
	return &OrderRepository{
		Collection: collection,
	}
}
func (r *OrderRepository) CreateOrder(ctx context.Context, order *models.Order) error {
	if order.Id == "" {
		order.Id = primitive.NewObjectID().Hex()
	}

	_, err := r.Collection.InsertOne(ctx, order)
	return err
}
