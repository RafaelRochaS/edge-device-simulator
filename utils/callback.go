package utils

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
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

	body, err := json.Marshal(data)

	if err != nil {
		log.Fatal("Failed to parse body: ", err)
	}

	log.Println("Body:", string(body))

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))

	if err != nil {
		log.Fatal("Failed to send callback: ", err)
	}

	log.Println("Response status:", resp.StatusCode)
	log.Println("Callback sent successfully.")
}
