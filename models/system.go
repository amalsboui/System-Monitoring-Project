package models

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