package kafkautil

import (
	"context"
	"fmt"
	"sync"

	"github.com/segmentio/kafka-go"
)

var (
	writers = map[string]*kafka.Writer{}
	mu      sync.Mutex
)

func NewWriter(topic string) *kafka.Writer {
	mu.Lock()
	defer mu.Unlock()

	if w, ok := writers[topic]; ok {
		return w
	}

	w := &kafka.Writer{
		Addr:         kafka.TCP("kafka:9092"),
		Topic:        topic,
		BatchSize:    1000,
		BatchTimeout: 500000000, // 500ms
		RequiredAcks: 1,
	}
	writers[topic] = w
	return w
}

func Publish(topic, msg string) error {
	w := NewWriter(topic)
	return w.WriteMessages(context.Background(), kafka.Message{Value: []byte(msg)})
}

func CloseWriters() {
	mu.Lock()
	defer mu.Unlock()
	for _, w := range writers {
		w.Close()
	}
}

func NewReader(topic, group string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{"kafka:9092"},
		Topic:          topic,
		GroupID:        group,
		StartOffset:    kafka.FirstOffset,
		CommitInterval: 0,
		MinBytes:       1,
		MaxBytes:       10e6,
	})
}

func ToCSV(id int, name, address, continent string) string {
	return fmt.Sprintf("%d,%s,%s,%s", id, name, address, continent)
}
