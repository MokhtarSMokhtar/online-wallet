package messsage

import (
	"encoding/json"
	md "github.com/MokhtarSMokhtar/online-wallet/comman/models"
	amqp "github.com/rabbitmq/amqp091-go"

	"log"
)

func (r *RabbitMQ) PublishUserRegisterEvent(event md.UserRegisteredEvent) error {

	body, err := json.Marshal(event)
	if err != nil {
		return err
	}
	err = r.Channel.Publish(
		"identity_exchange", // exchange
		"user.registered",   // routing key
		false,               // mandatory
		false,               // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		return err
	}

	log.Printf("Published UserRegisteredEvent: %+v", event)
	return nil

}
