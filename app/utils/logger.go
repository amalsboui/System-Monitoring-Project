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




func LogAlert(hostname, alertType string, value float64, unit string) {
    log.Printf("type=alert hostname=%s alert=%s value=%.2f unit=%s", 
        hostname, alertType, value, unit)
}