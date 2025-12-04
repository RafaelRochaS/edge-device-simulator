package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
)

func makePostCall(data any, url string) (string, error) {
	slog.Debug("Sending POST request to: ", slog.String("url", url))
	body, err := json.Marshal(data)

	if err != nil {
		slog.Error("Failed to parse body: ", slog.Any("error", err))
		return "", err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))

	if err != nil {
		slog.Error("Failed to send callback: ", slog.Any("error", err))
		return "", err
	}

	slog.Debug("Response status:", slog.Int("statusCode", resp.StatusCode))
	slog.Debug("Request sent successfully.")

	respBody, err := io.ReadAll(resp.Body)

	slog.Debug("Response body:", slog.String("body", string(respBody)))

	return string(respBody), err
}
