package esearch

import (
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
)

var client *elasticsearch.Client

func Init() {
	var err error
	URL := os.Getenv("ES_URL")
	userName := os.Getenv("ES_USERNAME")
	pwd := os.Getenv("ES_PASSWORD")
	cfg := elasticsearch.Config{
		Addresses: []string{
			URL,
		},
		Username: userName,
		Password: pwd,
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: 10 * time.Second,
			DialContext:           (&net.Dialer{Timeout: 10 * time.Second}).DialContext,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	client, err = elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the Elasticsearch client: %s", err)
	}
}

func GetClient() *elasticsearch.Client {
	return client
}
