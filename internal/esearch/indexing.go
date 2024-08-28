package esearch

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/esutil"
)

func LoadMapping(filename string) (string, error) {
	path := filepath.Join("internal", "mappings", filename)
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("error reading mapping file: %v", err)
	}
	return string(data), nil
}

func CheckIndexExists(indexName string) (bool, error) {
	client := GetClient()

	req := esapi.IndicesExistsRequest{
		Index: []string{indexName},
	}

	res, err := req.Do(context.Background(), client)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	if res.StatusCode == 200 {
		return true, nil
	}

	return false, nil
}

func CreateIndex(indexName, mappingFile string) error {
	client := GetClient()

	mapping, err := LoadMapping(mappingFile)
	if err != nil {
		return err
	}

	exists, err := CheckIndexExists(indexName)
	if err != nil {
		return err
	}

	if !exists {
		req := esapi.IndicesCreateRequest{
			Index: indexName,
			Body:  strings.NewReader(mapping),
		}

		res, err := req.Do(context.Background(), client)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		if res.IsError() {
			return fmt.Errorf("error creating index: %s", res.String())
		}

		log.Printf("Index %s created successfully.", indexName)
	} else {
		log.Printf("Index %s already exists.", indexName)
	}

	return nil
}

func UpdateIndexMapping(indexName string, updatedMapping string) error {
	client := GetClient()

	req := esapi.IndicesPutMappingRequest{
		Index: []string{indexName},
		Body:  strings.NewReader(updatedMapping),
	}

	res, err := req.Do(context.Background(), client)
	if err != nil {
		return fmt.Errorf("error updating index mapping: %v", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error response from Elasticsearch: %s", res.String())
	}

	log.Printf("Successfully updated mapping for index %s.", indexName)
	return nil
}

func IndexSingleDocument(indexName string, document string, documentID string) error {
	client := GetClient()

	req := esapi.IndexRequest{
		Index:      indexName,
		DocumentID: documentID,
		Body:       strings.NewReader(document),
	}

	res, err := req.Do(context.Background(), client)
	if err != nil {
		return fmt.Errorf("error indexing document: %v", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error response from Elasticsearch: %s", res.String())
	}

	log.Printf("Successfully indexed document %s into %s.", documentID, indexName)
	return nil
}

func BulkIndexDocuments(indexName string, documents []string) error {
	client := GetClient()

	bulkIndexer, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:  indexName,
		Client: client,
	})
	if err != nil {
		return fmt.Errorf("error creating bulk indexer: %v", err)
	}

	for _, doc := range documents {
		err := bulkIndexer.Add(context.Background(), esutil.BulkIndexerItem{
			Action:     "index",
			DocumentID: "", // Optionally set document ID
			Body:       strings.NewReader(doc),
		})
		if err != nil {
			return fmt.Errorf("error adding document to bulk indexer: %v", err)
		}
	}

	err = bulkIndexer.Close(context.Background())
	if err != nil {
		return fmt.Errorf("error closing bulk indexer: %v", err)
	}

	log.Printf("Successfully indexed %d documents into %s.", len(documents), indexName)
	return nil
}
