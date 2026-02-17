package kafka

import (
	"context"
	"log"

	kafkaGo "github.com/segmentio/kafka-go"
)

type kafkaClient struct {
	producer KafkaProducer
	consumer KafkaConsumer
}

func NewKafkaClient(producer KafkaProducer, consumer KafkaConsumer) *kafkaClient {
	return &kafkaClient{
		producer: producer,
		consumer: consumer,
	}
}

func NewKafkaProducerClient(producer KafkaProducer) *kafkaClient {
	return &kafkaClient{
		producer: producer,
	}
}

func NewKafkaConsumerClient(consumer KafkaConsumer) *kafkaClient {
	return &kafkaClient{
		consumer: consumer,
	}
}

type KafkaProducer interface {
	SendMessage(topic, brokerAddress, partitionKey string, message []byte) error
}

type KafkaConsumer interface {
	ReadMessage(topic, brokerAddress, consumerGroup string, handler func(m kafkaGo.Message) error) (err error)
}

func (k *kafkaClient) SendMessage(topic, brokerAddress, partitionKey string, message []byte) error {
	writer := kafkaGo.NewWriter(kafkaGo.WriterConfig{
		Brokers:  []string{brokerAddress},
		Topic:    topic,
		Balancer: &kafkaGo.LeastBytes{},
	})
	defer writer.Close()

	if err := writer.WriteMessages(context.Background(),
		kafkaGo.Message{
			Key:   []byte(partitionKey),
			Value: message,
		},
	); err != nil {
		return err
	}
	log.Printf("message to %v, %v sent\n", topic, partitionKey)
	return nil
}

func (k *kafkaClient) ReadMessage(topic, brokerAddress, consumerGroup string, handler func(m kafkaGo.Message) error) (err error) {
	handleMessage := func(ctx context.Context, m kafkaGo.Message, r *kafkaGo.Reader) {
		if err := handler(m); err != nil {
			log.Printf("handle message error: offset %v, partition %v, error: %v\n", m.Offset, m.Partition, err)
			return
		}
		if err := r.CommitMessages(ctx, m); err != nil {
			log.Printf("commit kafka message error:  offset %v, partition %v, error: %v\n", m.Offset, m.Partition, err)
		}
	}

	r := kafkaGo.NewReader(kafkaGo.ReaderConfig{
		Brokers:        []string{brokerAddress},
		Topic:          topic,
		GroupID:        consumerGroup,
		CommitInterval: 0,
	})
	defer r.Close()

	var m kafkaGo.Message
	for {
		ctx := context.Background()
		if m, err = r.ReadMessage(ctx); err != nil {
			return err
		}
		go handleMessage(ctx, m, r)
	}
}
