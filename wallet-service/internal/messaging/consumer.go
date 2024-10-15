package messaging

import (
	"encoding/json"
	md "github.com/MokhtarSMokhtar/online-wallet/comman/models"
	"github.com/MokhtarSMokhtar/online-wallet/wallet-service/internal/database"
	repository "github.com/MokhtarSMokhtar/online-wallet/wallet-service/internal/wallet-repository"
	"log"
)

// ConsumeUserRegisteredEvents starts consuming UserRegisteredEvent messages
func (r *RabbitMQ) ConsumeUserRegisteredEvents() {
	msgs, err := r.Channel.Consume(
		"wallet_queue", // queue
		"",             // consumer
		false,          // auto-ack
		false,          // exclusive
		false,          // no-local
		false,          // no-wait
		nil,            // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	go func() {
		for d := range msgs {
			var event md.UserRegisteredEvent
			if err := json.Unmarshal(d.Body, &event); err != nil {
				log.Printf("Error decoding UserRegisteredEvent: %v", err)
				err := d.Nack(false, false)
				if err != nil {
					return
				} // Reject the message without requeue
				continue
			}

			// Process the event
			log.Printf("Received UserRegisteredEvent: %+v", event)
			err := r.HandleUserRegisteredEvent(event)
			if err != nil {
				log.Printf("Failed to handle UserRegisteredEvent: %v", err)
				d.Nack(false, true) // Reject the message and requeue for retry
				continue
			}

			d.Ack(false) // Acknowledge the message upon successful processing
		}
	}()
}

// HandleUserRegisteredEvent processes the UserRegisteredEvent
func (r *RabbitMQ) HandleUserRegisteredEvent(event md.UserRegisteredEvent) error {
	// Implement the logic to create a wallet for the user
	db := database.GetDB()
	repo := repository.NewWalletRepository(db)
	err := repo.AddUserWallet(event.UserID)
	if err != nil {
		return err
	}

	log.Printf("Wallet created for UserID: %d", event.UserID)
	return nil
}
