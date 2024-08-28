package services

import (
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

type IndexingService interface {
	RegisterIndex() error
	ListenForMessages() error
}

type ServiceManager struct {
	services []IndexingService
	esClient *elasticsearch.Client
}

func NewServiceManager(esClient *elasticsearch.Client) *ServiceManager {
	sm := &ServiceManager{
		esClient: esClient,
	}
	sm.registerServices()
	return sm
}

func (sm *ServiceManager) registerServices() {
	productService := NewProductsService()
	sm.services = append(sm.services, productService)
}

func (sm *ServiceManager) StartAll() {
	for _, service := range sm.services {
		if err := service.RegisterIndex(); err != nil {
			log.Printf("Failed to register index: %v", err)
			continue
		}
		go func(s IndexingService) {
			if err := s.ListenForMessages(); err != nil {
				log.Printf("Failed to listen for messages: %v", err)
			}
		}(service)
	}
}
