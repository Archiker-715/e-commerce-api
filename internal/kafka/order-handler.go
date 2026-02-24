package kafka

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Archiker-715/e-commerce-api/internal/entity"
	kafkaGo "github.com/segmentio/kafka-go"
)

type paid interface {
	MarkPaid(orderId string) error
}

type orderHandler struct {
	usecase paid
	client  *kafkaConsumerClient
}

func NewKafkaOrderHandler(uc paid, client *kafkaConsumerClient) *orderHandler {
	return &orderHandler{
		usecase: uc,
		client:  client,
	}
}

func (o *orderHandler) Start() {
	defer o.client.reader.Close()
	for {
		ctx := context.Background()
		m, err := o.client.reader.ReadMessage(ctx)
		if err != nil {
			log.Printf("read kafka message error: key %v, partition %v. Err: %v\n", m.Key, m.Partition, err)
			continue
		}
		go o.handleOrderMessage(ctx, m)
	}
}

func (o *orderHandler) handleOrderMessage(ctx context.Context, m kafkaGo.Message) {
	var paidOrder entity.Paid
	if err := json.Unmarshal(m.Value, &paidOrder); err != nil {
		log.Printf("unmashal err kafka message: key %v, partition %v. Err: %v\n", m.Key, m.Partition, err)
		return
	}
	if err := o.usecase.MarkPaid(paidOrder.OrderId); err != nil {
		log.Printf("mark paid order error: offset %v, partition %v, error: %v\n", m.Offset, m.Partition, err)
		return
	}
	if err := o.client.reader.CommitMessages(ctx, m); err != nil {
		log.Printf("commit kafka message error:  offset %v, partition %v, error: %v\n", m.Offset, m.Partition, err)
		return
	}
}
