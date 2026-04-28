package domain

import "time"

type CPUMetrics struct {
	Usage float64
	Temp  float64
}

type GPUMetrics struct {
	Usage float64
	Temp  float64
}

type NetworkMetrics struct {
	UploadSpeed   float64
	DownloadSpeed float64
}

type MetricsSnapshot struct {
	CPU       CPUMetrics
	GPU       GPUMetrics
	Network   NetworkMetrics
	Timestamp time.Time
}
