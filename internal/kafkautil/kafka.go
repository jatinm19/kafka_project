package kafkautil

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func NewWriter(topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP("kafka:9092"),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

var writers = map[string]*kafka.Writer{}

func Publish(topic string, msg string) error {
	w, ok := writers[topic]

	if !ok {
		w = NewWriter(topic)
		writers[topic] = w
	}

	return w.WriteMessages(
		context.Background(),
		kafka.Message{
			Value: []byte(msg),
		},
	)
}

func CloseWriters() {
	for _, w := range writers {
		w.Close()
	}
}

func NewReader(topic, group string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{"kafka:9092"},
		Topic:       topic,
		GroupID:     group,
		StartOffset: kafka.FirstOffset,
		MinBytes:    1,
		MaxBytes:    10e6,
	})
}

func ToCSV(id int, name, address, continent string) string {
	return fmt.Sprintf("%d,%s,%s,%s", id, name, address, continent)
}
