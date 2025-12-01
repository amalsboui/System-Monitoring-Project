package kafkahelper

import (
	"context"
	"encoding/json"
	"log"

	"Event/models"
	"Event/utils"

	"github.com/segmentio/kafka-go"
)

func SendToKafka(writer *kafka.Writer, event models.SystemEvent){
	data, err := json.Marshal(event) //event struct to json
	if err!= nil{
		log.Println("Error marshaling event"+ err.Error())
		return
	}
	err = writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:  []byte(event.Hostname),
			Value: data,
		},
	)
	if err!= nil{
		log.Println("Error sending to Kafka"+ err.Error())
		return
	}
	utils.PrettyPrint(&event)
}