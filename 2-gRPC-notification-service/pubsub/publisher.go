package pubsub

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

func connectToRabbitMQ() (*amqp091.Connection, error) {
	rabbitMQHost := os.Getenv("RABBITMQ_HOST")
	var conn *amqp091.Connection
	var err error
	for i := 0; i < 5; i++ {
		conn, err = amqp091.Dial(fmt.Sprintf("amqp://guest:guest@%s:5672/", rabbitMQHost))
		if err == nil {
			return conn, nil
		}
		log.Printf("Failed to connect to RabbitMQ. Retrying in 2 seconds... (%d/5)", i+1)
		time.Sleep(2 * time.Second)
	}
	return nil, err
}

func Publish(eventType, filter, message string) {

	conn, err := connectToRabbitMQ()
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"notifications",
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare exchange: %v", err)
	}

	routingKey := eventType + "." + filter
	err = ch.Publish(
		"notifications",
		routingKey,
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		log.Fatalf("Failed to publish message: %v", err)
	}
}
