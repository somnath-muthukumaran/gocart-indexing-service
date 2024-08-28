package services

import (
	"log"

	"github.com/niklaus-mikael/gocart/indexing-service/internal/esearch"
)

type ProductsService struct {
	IndexingService
}

func NewProductsService() *ProductsService {
	return &ProductsService{}
}
func (p *ProductsService) RegisterIndex() error {
	indexName := "products"
	mappingFile := "mappings/product_mappings.json"
	err := esearch.CreateIndex(indexName, mappingFile)
	if err != nil {
		log.Printf("Failed to register index for %s: %v", indexName, err)
		return err
	}

	log.Printf("Successfully registered index for %s", indexName)
	return nil
}

func (p *ProductsService) ListenForMessages() error {
	log.Println("Listening for product messages")
	return nil
}
