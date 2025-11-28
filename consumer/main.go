package main

import (
	"log"
    "Event/kafkahelper"
	"Event/processor"

	"github.com/segmentio/kafka-go"
)


func main() {
	log.Println(("Starting Event Consumer..."))

	reader:= kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic: "system.metrics",
		GroupID: "metrics-consumer-group",
	})
	defer reader.Close()

	for{
		event := kafkahelper.ReadEvent(reader)
		processor.ProcessEvent(event)
	}
}