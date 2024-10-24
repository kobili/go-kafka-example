package kafka_client

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"log"
)

func ConnectProducer() *kafka.Producer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "kafka"})
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}

	log.Println("Created Kafka producer")

	return p
}
