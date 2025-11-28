package kafkahelper

import (
	"context"
	"encoding/json"
	"log"

	"Event/models"

	"github.com/segmentio/kafka-go"
)

//Read and Unmarshal kafka message into systemevent
func ReadEvent(reader *kafka.Reader) *models.SystemEvent {
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
