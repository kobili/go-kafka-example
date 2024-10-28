package kafka_client

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type KafkaClient struct {
	Producer *kafka.Producer
}

func NewProducer() (*KafkaClient, error) {
	producer, err := connectProducer()
	if err != nil {
		return nil, err
	}

	return &KafkaClient{
		Producer: producer,
	}, nil
}

func (k *KafkaClient) SendMessage(topic string, value []byte) error {
	fmt.Printf("Sending message to topic %s", topic)

	deliveryChan := make(chan kafka.Event)

	err := k.Producer.Produce(
		&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: kafka.PartitionAny,
			},
			Value: []byte(value),
		},
		deliveryChan,
	)

	if err != nil {
		return fmt.Errorf("failed to send message to kafka: %v", err)
	}

	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		return fmt.Errorf("failed to send message to kafka: %v", m.TopicPartition.Error)
	}

	close(deliveryChan)
	fmt.Printf("Sent message to topic %s", topic)
	return nil
}

func (k *KafkaClient) Close() {
	k.Producer.Close()
}
