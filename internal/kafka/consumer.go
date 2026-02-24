package kafka

import kafkaGo "github.com/segmentio/kafka-go"

type kafkaConsumerClient struct {
	reader *kafkaGo.Reader
}

type readers struct {
	Readers map[string]*kafkaGo.Reader
}

type KafkaConsumer interface {
	ReadMessage(handler func(m kafkaGo.Message) error) (err error)
}

func InitReader() *readers {
	return &readers{Readers: make(map[string]*kafkaGo.Reader)}
}

func NewKafkaConsumerClient(reader *kafkaGo.Reader) *kafkaConsumerClient {
	return &kafkaConsumerClient{
		reader: reader,
	}
}

func (r *readers) NewKafkaReader(topic, brokerAddress, consumerGroup string) *kafkaGo.Reader {
	newReader := kafkaGo.NewReader(kafkaGo.ReaderConfig{
		Brokers:        []string{brokerAddress},
		Topic:          topic,
		GroupID:        consumerGroup,
		CommitInterval: 0,
	})
	r.Readers[topic] = newReader
	return newReader
}
