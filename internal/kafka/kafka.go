package kafka

import (
	"context"
	"log"

	kafkaGo "github.com/segmentio/kafka-go"
)

type kafkaProducerClient struct {
	writer *kafkaGo.Writer
}

type kafkaConsumerClient struct {
	reader *kafkaGo.Reader
}

func NewKafkaProducerClient(writer *kafkaGo.Writer) *kafkaProducerClient {
	return &kafkaProducerClient{
		writer: writer,
	}
}

func NewKafkaWriter(topic, brokerAddress string) *kafkaGo.Writer {
	return kafkaGo.NewWriter(kafkaGo.WriterConfig{
		Brokers:  []string{brokerAddress},
		Topic:    topic,
		Balancer: &kafkaGo.LeastBytes{},
	})
}

func NewKafkaConsumerClient(reader *kafkaGo.Reader) *kafkaConsumerClient {
	return &kafkaConsumerClient{
		reader: reader,
	}
}

func NewKafkaReader(topic, brokerAddress, consumerGroup string) *kafkaGo.Reader {
	return kafkaGo.NewReader(kafkaGo.ReaderConfig{
		Brokers:        []string{brokerAddress},
		Topic:          topic,
		GroupID:        consumerGroup,
		CommitInterval: 0,
	})
}

type KafkaProducer interface {
	SendMessage(partitionKey string, message []byte) error
}

type KafkaConsumer interface {
	ReadMessage(handler func(m kafkaGo.Message) error) (err error)
}

func (k *kafkaProducerClient) SendMessage(partitionKey string, message []byte) error {
	defer k.writer.Close()

	if err := k.writer.WriteMessages(context.Background(),
		kafkaGo.Message{
			Key:   []byte(partitionKey),
			Value: message,
		},
	); err != nil {
		return err
	}
	log.Printf("message to %v, %v sent\n", k.writer.Topic, partitionKey)
	return nil
}

func (k *kafkaConsumerClient) ReadMessage(handler func(m kafkaGo.Message) error) (err error) {
	handleMessage := func(ctx context.Context, m kafkaGo.Message, r *kafkaGo.Reader) {
		if err := handler(m); err != nil {
			log.Printf("handle message error: offset %v, partition %v, error: %v\n", m.Offset, m.Partition, err)
			return
		}
		if err := r.CommitMessages(ctx, m); err != nil {
			log.Printf("commit kafka message error:  offset %v, partition %v, error: %v\n", m.Offset, m.Partition, err)
		}
	}

	defer k.reader.Close()

	var m kafkaGo.Message
	for {
		ctx := context.Background()
		if m, err = k.reader.ReadMessage(ctx); err != nil {
			return err
		}
		go handleMessage(ctx, m, k.reader)
	}
}
