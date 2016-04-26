package main

import (
	"encoding/json"
	"log"
	"os"
	"net/http"
    "bytes"
)

type HttpPublisher struct {
	uri string
}

func ConnectHttpPublisher() (Publisher, error) {
	uri := os.Getenv("HTTP_URI")
	publisher := &HttpPublisher{uri}
	log.Printf("Will HTTP POST to %s", uri)

	return publisher, nil
}

func (p *HttpPublisher) Configure() error {
	return nil
}

func (p *HttpPublisher) Publish(data interface{}) error {
	log.Printf("Publishing...")
	body, err := json.Marshal(data)
	if err != nil {
		log.Printf("Failed to marshal data: %s", err.Error())
		return err
	}

    req, err := http.NewRequest("POST", p.uri, bytes.NewBuffer(body))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

	return nil
}
