package main

import "Event/models"

func UpdateMetrics(event *models.SystemEvent) {
	if event == nil {
		return
	}

	totalCPU := event.CPUUser + event.CPUSystem

	cpuUsage.WithLabelValues(event.Hostname).Set(totalCPU)
	memoryUsage.WithLabelValues(event.Hostname).Set(event.MemoryUsed)
	diskUsage.WithLabelValues(event.Hostname).Set(event.DiskUsed)
	diskIO.WithLabelValues(event.Hostname).Set(event.DiskIO)
	netIn.WithLabelValues(event.Hostname).Set(event.NetIn)
	netOut.WithLabelValues(event.Hostname).Set(event.NetOut)














}