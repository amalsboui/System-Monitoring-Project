package main

import (
	"context"
	"fmt"
	"encoding/json"
	"log"
    "Event/models"
	"Event/utils"

	"github.com/segmentio/kafka-go"
)


//Read and Unmarshal kafka message into systemevent
func readEvent(reader *kafka.Reader) *models.SystemEvent {
	msg, err := reader.ReadMessage(context.Background())
	if err != nil {
		log.Println("Error reading message:" + err.Error())
		return nil
	}

	var event models.SystemEvent
	err = json.Unmarshal(msg.Value, &event)
	if err != nil{
		log.Println("Error unmarshaling message:" + err.Error())
		return nil
	}

	return &event
}

//handles the event(log it, filters it, and pretty print)
func processEvent(event *models.SystemEvent) {
	if event == nil {
		return
	}

	totalCPU := event.CPUUser + event.CPUSystem

	// CPU alert
	if totalCPU > 80 {
		utils.LogAlert(event.Hostname + " CPU high: " + formatFloat(totalCPU) + "%")
	}
	// Memory alert
	if event.MemoryUsed > 10 {
		utils.LogAlert(event.Hostname + " Memory high: " + formatFloat(event.MemoryUsed) + " GB")
	}
	// Disk alert
	if event.DiskUsed > 180 {
		utils.LogAlert(event.Hostname + " Disk high: " + formatFloat(event.DiskUsed) + " GB")
	}
	// Disk IO alert
	if event.DiskIO > 40 {
		utils.LogAlert(event.Hostname + " Disk I/O high: " + formatFloat(event.DiskIO) + " MB/s")
	}
	// Network alert
	if event.NetIn > 30 || event.NetOut > 30 {
		utils.LogAlert(event.Hostname + " Network high: In " + formatFloat(event.NetIn) + " MB/s, Out " + formatFloat(event.NetOut) + " MB/s")
	}

	utils.PrettyPrint(event)
}

func formatFloat(val float64) string {
	return fmt.Sprintf("%.2f", val)
}



func main() {
	log.Println(("Starting Event Consumer..."))

	reader:= kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic: "system.metrics",
		GroupID: "metrics-consumer-group",
	})
	defer reader.Close()

	for{
		event := readEvent(reader)
		processEvent(event)
	}
}