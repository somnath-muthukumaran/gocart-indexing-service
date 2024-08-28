package main

import (
	"fmt"
	"log"

	"github.com/niklaus-mikael/gocart/indexing-service/internal/config"
	"github.com/niklaus-mikael/gocart/indexing-service/internal/esearch"
	"github.com/niklaus-mikael/gocart/indexing-service/internal/rmq"
	"github.com/niklaus-mikael/gocart/indexing-service/internal/services"
)

func main() {
	fmt.Println("Checking")
	config.LoadConfig()
	esearch.Init()
	_, err := rmq.Connect()
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ:", err)
	}
	conn, err := rmq.GetConnection()
	if err != nil {
		log.Fatal("Failed to retrieve RabbitMQ connection:", err)
	}
	defer conn.Close()
	serviceManager := services.NewServiceManager()
	serviceManager.StartAll()

	forever := make(chan bool)
	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
