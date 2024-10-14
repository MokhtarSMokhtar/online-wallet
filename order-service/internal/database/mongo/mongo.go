package mongo

import (
	"context"
	"fmt"
	"github.com/MokhtarSMokhtar/online-wallet/order-service/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type OrderDB struct {
	*mongo.Client
}

func (db *OrderDB) ConnectToDb() error {
	cfg := config.NewConfig()

	var mongoURI string
	if cfg.DBUser != "" && cfg.DBPassword != "" {
		mongoURI = fmt.Sprintf("mongodb://%s:%s@%s:%s", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort)
	} else {
		mongoURI = fmt.Sprintf("mongodb://%s:%s", cfg.DBHost, cfg.DBPort)
	}
	fmt.Println(mongoURI)
	clientOptions := options.Client().ApplyURI(mongoURI)
	// Create a client
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return fmt.Errorf("failed to create new MongoDB client: %v", err)
	}

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect the client
	err = client.Connect(ctx)
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	// Ping the database to verify connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	fmt.Println("Connected to MongoDB!")
	db.Client = client
	return nil
}
