package services

import (
	"encoding/json"
	"log"

	"github.com/niklaus-mikael/gocart/indexing-service/internal/esearch"
	"github.com/niklaus-mikael/gocart/indexing-service/internal/rmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

type ProductsService struct {
	IndexingService
}

func NewProductsService() *ProductsService {
	return &ProductsService{}
}
func (p *ProductsService) RegisterIndex() error {
	indexName := "products"
	mappingFile := "product_mappings.json"
	err := esearch.CreateIndex(indexName, mappingFile)
	if err != nil {
		log.Printf("Failed to register index for %s: %v", indexName, err)
		return err
	}

	log.Printf("Successfully registered index for %s", indexName)
	return nil
}

func (p *ProductsService) ListenForMessages() error {
	err := rmq.Consume("product_queue", eventHandler)
	if err != nil {
		log.Fatalf("Error starting product service listener: %v", err)
	}
	return nil
}

func eventHandler(msg amqp.Delivery) {
	var productData map[string]interface{}
	if err := json.Unmarshal(msg.Body, &productData); err != nil {
		log.Printf("Error decoding product event: %v", err)
		return
	}
	log.Printf("Processing product event: %v", productData)
	// logic to be added here
}
