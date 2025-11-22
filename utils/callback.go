package utils

import (
	"log"
	"time"

	"github.com/RafaelRochaS/edge-device-simulator/models"
	"github.com/google/uuid"
)

func GetCallbackData() models.CallbackData {
	return models.CallbackData{
		TaskID:        uuid.New().String(),
		ExecutionSite: "local",
		CreatedAt:     time.Now().Unix(),
	}
}

func SendCallback(data models.CallbackData, url string) {
	log.Printf("Sending callback: %+v\n", data)
	_, err := makePostCall(data, url)

	if err != nil {
		log.Println("Failed to send callback: ", err)
	}
}
