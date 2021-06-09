package main

import (
	"context"
	"fmt"
	elasticsearch "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"

	"log"
	"strings"
)

//// Indexed task
//type RegionalMachine struct {
//	MachineTypes []string `json:"machine_types"`
//	Regions      []string `json:"regions"`
//	SnapshotTime string   `json:"snapshot_time"`
//}
//// Create a mapping for the Elasticsearch documents
//var (
//	docMap map[string]interface{}
//)
func main() {

	cfg := elasticsearch.Config{
		Addresses: []string{
			"machine ip with port",
		},
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	log.Println(res)

	fmt.Println("===================================================")
	res1, err := es.Index(
		"test",                                  // Index name
		strings.NewReader(`{"title" : "Test"}`), // Document body
		es.Index.WithDocumentID("1"),            // Document ID
		es.Index.WithRefresh("true"),            // Refresh
	)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	defer res1.Body.Close()

	log.Println(res1)

	fmt.Println("===================================================")

	req := esapi.IndexRequest{
		Index:      "test",                                  // Index name
		Body:       strings.NewReader(`{"title" : "Test"}`), // Document body
		DocumentID: "1",                                     // Document ID
		Refresh:    "true",                                  // Refresh
	}

	res2, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res2.Body.Close()

	log.Println(res2)

}

// json return data by server
