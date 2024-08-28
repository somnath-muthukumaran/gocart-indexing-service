package main

import (
	"fmt"

	"github.com/niklaus-mikael/gocart/indexing-service/internal/config"
	"github.com/niklaus-mikael/gocart/indexing-service/internal/esearch"
	"github.com/niklaus-mikael/gocart/indexing-service/internal/services"
)

func main() {
	fmt.Println("Checking")
	config.LoadEnv()
	esearch.Init()
	esClient := esearch.GetClient()
	serviceManager := services.NewServiceManager(esClient)
	serviceManager.StartAll()
}
