package processor

import (
	"fmt"

	"Event/models"
	"Event/utils"
)


//handles the event(log it, filters it, and pretty print)
func ProcessEvent(event *models.SystemEvent) {
	if event == nil {
		return
	}

	totalCPU := event.CPUUser + event.CPUSystem

	// CPU alert
	if totalCPU > 80 {
		utils.LogAlert(event.Hostname + " CPU high: " + fmt.Sprintf("%.2f", totalCPU) + "%")
	}
	// Memory alert
	if event.MemoryUsed > 10 {
		utils.LogAlert(event.Hostname + " Memory high: " + fmt.Sprintf("%.2f", event.MemoryUsed) + " GB")
	}
	// Disk alert
	if event.DiskUsed > 180 {
		utils.LogAlert(event.Hostname + " Disk high: " + fmt.Sprintf("%.2f", event.DiskUsed) + " GB")
	}
	// Disk IO alert
	if event.DiskIO > 40 {
		utils.LogAlert(event.Hostname + " Disk I/O high: " + fmt.Sprintf("%.2f", event.DiskIO) + " MB/s")
	}
	// Network alert
	if event.NetIn > 30 || event.NetOut > 30 {
		utils.LogAlert(event.Hostname + " Network high: In " + fmt.Sprintf("%.2f", event.NetIn) +
			" MB/s, Out " + fmt.Sprintf("%.2f", event.NetOut) + " MB/s")
	}

	utils.PrettyPrint(event)
}
