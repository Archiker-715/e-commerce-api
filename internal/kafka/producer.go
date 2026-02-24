package kafka

import kafkaGo "github.com/segmentio/kafka-go"

type kafkaProducerClient struct {
	writer *kafkaGo.Writer
}

type writers struct {
	Writers map[string]*kafkaGo.Writer
}

type KafkaProducer interface {
	SendMessage(partitionKey string, message []byte) error
}

func InitWriter() *writers {
	return &writers{Writers: make(map[string]*kafkaGo.Writer)}
}

func NewKafkaProducerClient(writer *kafkaGo.Writer) *kafkaProducerClient {
	return &kafkaProducerClient{
		writer: writer,
	}
}

func (w *writers) NewKafkaWriter(topic, brokerAddress string) *kafkaGo.Writer {
	newWriter := kafkaGo.NewWriter(kafkaGo.WriterConfig{
		Brokers:  []string{brokerAddress},
		Topic:    topic,
		Balancer: &kafkaGo.LeastBytes{},
	})
	w.Writers[topic] = newWriter
	return newWriter
}
