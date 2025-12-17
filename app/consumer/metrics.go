package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var(
	cpuUsage = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "system_cpu_usage_percent",
			Help: "CPU usage percentage",
		},
		[]string{"hostname"},//to let grafana split lines epr server
	)

	memoryUsage = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "system_memory_usage_gb",
			Help: "Memory usage in GB",
		},
		[]string{"hostname"},
	)

	diskUsage = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "system_disk_usage_gb",
			Help: "Disk usage in GB",
		},
		[]string{"hostname"},
	)

	diskIO = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "system_disk_io_mb_per_sec",
			Help: "Disk IO in MB/s",
		},
		[]string{"hostname"},
	)

	netIn = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "system_network_in_mb_per_sec",
			Help: "Network inbound traffic",
		},
		[]string{"hostname"},
	)

	netOut = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "system_network_out_mb_per_sec",
			Help: "Network outbound traffic",
		},
		[]string{"hostname"},

	)
)

func InitMetrics(){
	prometheus.MustRegister(cpuUsage)
	prometheus.MustRegister(memoryUsage)
	prometheus.MustRegister(diskUsage)
	prometheus.MustRegister(diskIO)
	prometheus.MustRegister(netIn)
	prometheus.MustRegister(netOut)
}

func StartMetricsServer() {
	http.Handle("/metrics", promhttp.Handler())
    go http.ListenAndServe(":2112", nil)
}