package main

import (
	"log"

	"github.com/nicolasjhampton/kafka-demo/cmd/producer"
)

func main() {
	producer, err := producer.Setup()
	if err != nil {
		log.Fatalf("failed to init producer: %v", err)
	}
	defer producer.Close()
}
