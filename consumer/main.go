package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

type SystemEvent struct{
	Hostname     string  `json:"hostname"`
	CPUUser      float64 `json:"cpu_user"`
	CPUSystem    float64 `json:"cpu_system"`
	CPUIdle      float64 `json:"cpu_idle"`
	MemoryUsed   float64 `json:"memory_used_gb"`
	MemoryFree   float64 `json:"memory_free_gb"`
	MemoryCache  float64 `json:"memory_cache_gb"`
	DiskUsed     float64 `json:"disk_used_gb"`
	DiskFree     float64 `json:"disk_free_gb"`
	DiskIO       float64 `json:"disk_io_mb_s"`
	NetIn        float64 `json:"net_in_mb_s"`
	NetOut       float64 `json:"net_out_mb_s"`
	Timestamp    string  `json:"timestamp"`
}

//Read and Unmarshal kafka message into systemevent
func readEvent(reader *kafka.Reader) *SystemEvent {
	msg, err := reader.ReadMessage(context.Background())
	if err != nil {
		log.Println("Error reading message:", err)
		return nil
	}

	var event SystemEvent
	err = json.Unmarshal(msg.Value, &event)
	if err != nil{
		log.Println("Error unmarshaling message:", err)
		return nil
	}

	return &event
}

//handles the event(log it, filters it, and pretty print)
func processEvent(event *SystemEvent) {
	if event == nil {
		return
	}

	totalCPU := event.CPUUser + event.CPUSystem

	prettyJSON, _ := json.MarshalIndent(event, "", "  ")
	fmt.Println(string(prettyJSON))

	// CPU alert
    if totalCPU > 80 {
        log.Printf("High CPU Alert! %s: %.2f%% used\n", event.Hostname, totalCPU)
    }

    // Memory alert
    if event.MemoryUsed > 10 {
        log.Printf("High Memory Alert! %s: %.2f GB used\n", event.Hostname, event.MemoryUsed)
    }

    // Disk alert
    if event.DiskUsed > 180 {
        log.Printf("High Disk Usage! %s: %.2f GB used\n", event.Hostname, event.DiskUsed)
    }

    // Disk IO alert
    if event.DiskIO > 40 {
        log.Printf("High Disk I/O! %s: %.2f MB/s\n", event.Hostname, event.DiskIO)
    }

    // Network alert
    if event.NetIn > 30 || event.NetOut > 30 {
        log.Printf("High Network Traffic! %s: In: %.2f Out: %.2f MB/s\n",
            event.Hostname, event.NetIn, event.NetOut)
    }
}



func main() {
	fmt.Println(("Starting Event Consumer..."))

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