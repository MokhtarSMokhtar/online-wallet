package messsage

import (
	"fmt"
	"github.com/MokhtarSMokhtar/online-wallet/comman/config"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"sync"
)

type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

var (
	instance *RabbitMQ
	once     sync.Once
)

// GetRabbitMQInstance initializes and returns the singleton RabbitMQ instance
func GetRabbitMQInstance() *RabbitMQ {
	once.Do(func() {
		cfg := config.NewConfig()
		connStr := fmt.Sprintf("amqp://%s:%s@%s:%s/",
			cfg.RabbitMQUser,
			cfg.RabbitMQPassword,
			cfg.RabbitMQHost,
			cfg.RabbitMQPort,
		)

		conn, err := amqp.Dial(connStr)
		if err != nil {
			log.Fatalf("Failed to connect to RabbitMQ: %v", err)
		}

		ch, err := conn.Channel()
		if err != nil {
			log.Fatalf("Failed to open a channel: %v", err)
		}

		// Declare exchange
		err = ch.ExchangeDeclare(
			"identity_exchange", // name
			"direct",            // type
			true,                // durable
			false,               // auto-deleted
			false,               // internal
			false,               // no-wait
			nil,                 // arguments
		)
		if err != nil {
			log.Fatalf("Failed to declare exchange: %v", err)
		}

		instance = &RabbitMQ{
			Conn:    conn,
			Channel: ch,
		}

		log.Println("Connected to RabbitMQ")
	})

	return instance
}

// Close closes the RabbitMQ connection and channel
func (r *RabbitMQ) Close() {
	if r.Channel != nil {
		r.Channel.Close()
	}
	if r.Conn != nil {
		r.Conn.Close()
	}
}
