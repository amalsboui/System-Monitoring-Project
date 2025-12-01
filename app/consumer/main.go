package main

import (
	"log"
    "Event/kafkahelper"
	"Event/processor"

	"github.com/segmentio/kafka-go"
	"os"
)


func main() {
	file, err := os.OpenFile("/app/logs/consumer.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)

	log.Println(("Starting Event Consumer..."))

	reader:= kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"kafka:9092"},
		Topic: "system.metrics",
		GroupID: "metrics-consumer-group",
	})
	defer reader.Close()

	for{
		event := kafkahelper.ReadEvent(reader)
		processor.ProcessEvent(event)
	}
}