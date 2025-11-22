package models

type Task struct {
	Id          string `json:"id"`
	DeviceId    string `json:"deviceId"`
	Image       string `json:"image"`
	CPU         string `json:"cpu"`
	Mem         int    `json:"mem"`
	CallbackUrl string `json:"callbackUrl"`
	Workload    int    `json:"workload"`
}

type StartTaskRequest struct {
	Id       string `json:"id"`
	DeviceId string `json:"deviceId"`
}

type RegisterTaskResponse struct {
	Id string `json:"id"`
}
