package database

import (
	"context"
	"fmt"
	"github.com/MokhtarSMokhtar/online-wallet/payment-service/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"sync"
	"time"
)

type MongoDB struct {
	Client   *mongo.Client
	Database *mongo.Database
}

var (
	instance *MongoDB
	once     sync.Once
)

func GetMongoClient(ctx context.Context) *MongoDB {

	once.Do(func() {
		cfg := config.NewConfig()
		var mongoURI string
		if cfg.DBUser != "" && cfg.DBPassword != "" {
			mongoURI = fmt.Sprintf("mongodb://%s:%s@%s:%s", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort)
		} else {
			mongoURI = fmt.Sprintf("mongodb://%s:%s", cfg.DBHost, cfg.DBPort)
		}

		clientOptions := options.Client().ApplyURI(mongoURI)
		clientOptions.SetConnectTimeout(10 * time.Second) // Setting a timeout for connection

		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Fatalf("Failed to create MongoDB client: %v", err)
		}

		// Ping the database to verify the connection
		err = client.Ping(ctx, nil)
		if err != nil {
			log.Fatalf("Failed to connect to MongoDB: %v", err)
		}

		db := client.Database(cfg.DBName)

		instance = &MongoDB{
			Client:   client,
			Database: db,
		}

		log.Println("Successfully connected to MongoDB!")
	})

	return instance
}

func (m *MongoDB) Close(ctx context.Context) {
	if err := m.Client.Disconnect(ctx); err != nil {
		log.Printf("Error disconnecting from MongoDB: %v", err)
	} else {
		log.Println("Successfully disconnected from MongoDB.")
	}
}
