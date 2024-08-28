package rmq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Consume(queueName string, handler func(amqp.Delivery)) error {
	// Ensure the connection is established
	conn, err := GetConnection()
	if err != nil {
		return err
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
		return err
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
		return err
	}

	forever := make(chan bool)

	go func() {
		for msg := range msgs {
			handler(msg) // Pass the message to the provided handler function
		}
	}()

	log.Printf("Waiting for messages on queue: %s", queueName)
	<-forever

	return nil
}
