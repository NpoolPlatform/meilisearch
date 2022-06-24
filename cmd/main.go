package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/meilisearch/meilisearch-go"
)

func main() {
	log.SetFlags(log.Lshortfile)

	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host: "http://127.0.0.1:7700",
	})

	jsonFile, err := os.Open("./movies.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	var movies []map[string]interface{}
	if err := json.Unmarshal(byteValue, &movies); err != nil {
		log.Fatal(err)
	}

	_, err = client.Index("movies").AddDocuments(movies)
	if err != nil {
		log.Fatal(err)
	}
}
