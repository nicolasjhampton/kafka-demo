package main

import (
	"log"

	"github.com/nicolasjhampton/kafka-demo/cmd/producer"
	routes "github.com/nicolasjhampton/kafka-demo/internal/api/http"
)

// func setupRouter() *gin.Engine {
// 	r := gin.Default()

// 	return r
// }

func main() {
	producer, err := producer.Setup()
	if err != nil {
		log.Fatalf("failed to init producer: %v", err)
	}
	defer producer.Close()

	r := routes.SetupRouter()

	// routes
	routes.PingRoute(r)

	if err := r.Run(); err != nil {
		log.Printf("failed to run the server: %v", err)
	}
}
