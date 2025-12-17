package processor

import (
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
		utils.LogAlert(event.Hostname, "cpu", totalCPU, "%")
	}
	// Memory alert
	if event.MemoryUsed > 10 {
		utils.LogAlert(event.Hostname, "memory", event.MemoryUsed, "GB")
	}
	// Disk alert
	if event.DiskUsed > 180 {
		utils.LogAlert(event.Hostname, "disk", event.DiskUsed, "GB")
	}
	// Disk IO alert
	if event.DiskIO > 40 {
		utils.LogAlert(event.Hostname, "disk_io", event.DiskIO, "MB/s")
	}
	// Network alert - log both separately
	if event.NetIn > 30 {
		utils.LogAlert(event.Hostname, "net_in", event.NetIn, "MB/s")
	}
	if event.NetOut > 30 {
		utils.LogAlert(event.Hostname, "net_out", event.NetOut, "MB/s")
	}

	utils.PrettyPrint(event)
}
