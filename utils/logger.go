package utils

import (
	"encoding/json"
	"fmt"
	"log"

	"Event/models"
)

// PrettyPrint prints event as indented JSON
func PrettyPrint(event *models.SystemEvent) {
	if event == nil {
		return
	}
	data, _ := json.MarshalIndent(event, "", "  ")
	fmt.Println(string(data))
}

// LogAlert prints a log alert
func LogAlert(msg string) {
	log.Println("⚠️ ", msg)
}