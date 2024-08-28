package rmq

import (
	"log"
	"time"

	"github.com/niklaus-mikael/gocart/indexing-service/internal/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

var conn *amqp.Connection
var err error

func Connect() (*amqp.Connection, error) {
	cnf := config.GetEnvDetails()
	for i := 0; i < 5; i++ {
		conn, err = amqp.Dial(cnf.RMQ_URL)
		if err == nil {
			log.Println("Connected to RabbitMQ successfully")
			return conn, nil
		}
		log.Printf("Failed to connect to RabbitMQ (attempt %d): %v", i+1, err)
		time.Sleep(2 * time.Second)
	}

	return nil, err
}
func GetConnection() (*amqp.Connection, error) {
	if conn == nil {
		return nil, err
	}
	return conn, nil
}
