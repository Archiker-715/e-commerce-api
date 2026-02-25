package app

import "github.com/Archiker-715/e-commerce-api/internal/kafka"

func (a *app) startKafkaConsumers() {
	kafkaHandler := kafka.NewKafkaOrderHandler(
		a.Services.OrderService,
		kafka.NewKafkaConsumerClient(
			kafka.InitReader().NewKafkaReader("paid-orders", "localhost", "test-group"),
		),
	)
	kafkaHandler.Start()
}
