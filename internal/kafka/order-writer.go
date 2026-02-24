package kafka

import (
	"context"
	"encoding/json"
	"log"

	ctxpkg "github.com/Archiker-715/e-commerce-api/internal/auth/ctx"
	"github.com/Archiker-715/e-commerce-api/internal/entity"
	kafkaGo "github.com/segmentio/kafka-go"
)

type PaidOrderEventPublisher interface {
	PaidEventPublish(ctx context.Context, orderId string)
}

type OrderWriter struct {
	client *kafkaProducerClient
}

func NewKafkaOrderWriter(client *kafkaProducerClient) *OrderWriter {
	return &OrderWriter{
		client: client,
	}
}

func (o *OrderWriter) PaidEventPublish(ctx context.Context, orderId string) {
	defer o.client.writer.Close()

	jsonB, err := json.Marshal(entity.Paid{OrderId: orderId})
	if err != nil {
		log.Printf("marshal orderId err: %v\n", err)
	}

	partitionKey := ctxpkg.UserFromCtxAsStr(ctx)
	if err := o.client.writer.WriteMessages(context.Background(),
		kafkaGo.Message{
			Key:   []byte(partitionKey),
			Value: jsonB,
		},
	); err != nil {
		log.Printf("message write to %v, %v error: %v\n", o.client.writer.Topic, partitionKey, err)
	}
	log.Printf("message to %v, %v sent\n", o.client.writer.Topic, partitionKey)
}
