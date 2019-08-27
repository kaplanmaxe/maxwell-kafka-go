package main

import (
	"context"
	"flag"
	"fmt"

	kafka "github.com/segmentio/kafka-go"
)

func main() {
	var (
		bootstrapAddressFlag string
		topicFlag            string
		partitionFlag        int
	)

	flag.StringVar(&bootstrapAddressFlag, "bootstrapAddress", "127.0.0.1:9092", "address to bootstrap server")
	flag.StringVar(&topicFlag, "topic", "maxwell", "topic to subscribe to")
	flag.IntVar(&partitionFlag, "partition", 0, "partition number")
	flag.Parse()

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{fmt.Sprintf("%s", bootstrapAddressFlag)},
		Topic:     topicFlag,
		Partition: partitionFlag,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})
	r.SetOffset(0)

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}

	r.Close()
}
