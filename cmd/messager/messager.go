package producer

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
	routes "github.com/nicolasjhampton/kafka-demo/internal/api/http"
)

const kafkaServerAddress = "localhost:9092"

func Setup(address string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer([]string{address}, config)
	if err != nil {
		return nil, fmt.Errorf("failed to setup producer: %w", err)
	}
	return producer, nil
}

func main() {
	producer, err := Setup(kafkaServerAddress)
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
