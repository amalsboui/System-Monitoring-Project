package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/segmentio/kafka-go"
)

//Event structure
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

var servers = []string{"server1", "server2", "server3"}

func randomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
	//to generate random float
}

func generateEvent() SystemEvent {//systemevent houa return type
	cpuUser := randomFloat(10,80)
	cpuSystem := randomFloat(5,30)
	cpuIdle := 100 - cpuUser -cpuSystem

	return SystemEvent{
		Hostname: servers[rand.Intn(len(servers))],
		CPUUser: cpuUser,
		CPUSystem: cpuSystem,
		CPUIdle: cpuIdle,
		MemoryUsed:  randomFloat(4, 12),
		MemoryFree:  randomFloat(2, 10),
		MemoryCache: randomFloat(1, 4),
		DiskUsed:    randomFloat(80, 200),
		DiskFree:    randomFloat(20, 100),
		DiskIO:      randomFloat(5, 50),
		NetIn:       randomFloat(1, 40),
		NetOut:      randomFloat(1, 40),
		Timestamp:   time.Now().UTC().Format(time.RFC3339),
	}
}

func sendToKafka(writer *kafka.Writer, event SystemEvent){
	data, err := json.Marshal(event) //event struct to json
	if err!= nil{
		log.Println("Error marshaling event", err)
		return
	}
	err = writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:  []byte(event.Hostname),
			Value: data,
		},
	)
	if err!= nil{
		log.Println("Error sending to Kafka", err)
		return
	}
	prettyJSON, _ := json.MarshalIndent(event, "", "  ")
    log.Println(string(prettyJSON))
}

func main(){
	fmt.Println("Starting Event Generator...")

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic: "system.metrics",
	})
	defer writer.Close()

	for{
		event := generateEvent()
		sendToKafka(writer, event)
		time.Sleep(1 * time.Second)
	}
}