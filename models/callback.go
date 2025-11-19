package models

type CallbackData struct {
	TaskID        string  `json:"taskId"`
	DeviceID      int     `json:"deviceId"`
	WorkloadSize  int     `json:"workloadSize"`
	ExecutionSite string  `json:"executionSite"`
	CreatedAt     int64   `json:"createdAt"`
	Duration      float64 `json:"duration"`
}
