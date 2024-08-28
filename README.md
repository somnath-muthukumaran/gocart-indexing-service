# Go Fiber Elasticsearch Indexing Service

This project is a microservice built using [Go Fiber](https://gofiber.io/) that listens to RabbitMQ queues and indexes data into Elasticsearch. The service is designed to manage multiple indices, each with its own service for registration and message consumption.

## Features

- **Service-Oriented Design**: Each index has its own service for handling registration and message listening.
- **Elasticsearch Integration**: Automatically creates indices with predefined mappings if they don't exist.
- **RabbitMQ Integration**: Listens to queues for incoming data to be indexed into Elasticsearch.
- **Centralized Configuration**: Environment variables are loaded once and managed through a configuration struct.
- **Logging**: Unified logging using the custom logger.

## Project Structure

├── cmd/
│   └── main.go             # Entry point of the application
├── config/
│   └── config.go           # Configuration management
├── internal/
│   ├── esearch/            # Elasticsearch utility functions
|   |   ├── client.go       # Elasticsearch client setup and configuration
|   |   ├── indexing.go     # Index creation, updating, and document management
│   └── utils/
│       └── logger.go       # Custom logging utility
├── mappings/               # Directory for Elasticsearch index mappings
│   └── product_mappings.json  # Mapping for the "products" index
├── services/
│   ├── service_manager.go  # ServiceManager to manage and start services
│   ├── products.go         # Service handling "products" index
├── rmq/                    # RabbitMQ utilities
│   ├── connection.go       # RabbitMQ connection management
│   └── consumer.go         # RabbitMQ consumer for message handling
└── README.md               # This file

## Setup


1. **Clone the repository:**
    git clone https://github.com/somnath-muthukumaran/gocart-indexing-service.git
   
    cd your-repo-directory

3. **Install dependencies:**
    go mod tidy

4. **Set up environment variables:**
   Create a .env file in the root directory and define your environment variables (e.g., RabbitMQ and Elasticsearch configurations).

5. **Run the service:**
    go run cmd/main.go

## Contributing

Feel free to submit issues, fork the repository, and send pull requests.
