package utils

import (
	"log/slog"
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
	slog.Debug("Sending callback:", slog.Any("callback", data), slog.String("url", url))
	_, err := makePostCall(data, url)

	if err != nil {
		slog.Warn("Failed to send callback: ", slog.Any("error", err))
	}
}
