package main

import (
	"log"
	"math/rand"
	"time"

	"Event/kafkahelper"
	"Event/models"

	"github.com/segmentio/kafka-go"
)



var servers = []string{"server1", "server2", "server3"}

func randomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
	//to generate random float
}

func generateEvent() models.SystemEvent {//systemevent houa return type
	cpuUser := randomFloat(10,80)
	cpuSystem := randomFloat(5,30)
	cpuIdle := 100 - cpuUser -cpuSystem

	return models.SystemEvent{
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


func main(){
	log.Println("Starting Event Generator...")

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic: "system.metrics",
	})
	defer writer.Close()

	for{
		event := generateEvent()
		kafkahelper.SendToKafka(writer, event)
		time.Sleep(1 * time.Second)
	}
}