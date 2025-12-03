package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
)

func makePostCall(data any, url string) (string, error) {
	slog.Debug("Sending POST request to: ", url)
	body, err := json.Marshal(data)

	if err != nil {
		slog.Error("Failed to parse body: ", err)
		return "", err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))

	if err != nil {
		slog.Error("Failed to send callback: ", err)
		return "", err
	}

	slog.Debug("Response status:", resp.StatusCode)
	slog.Debug("Request sent successfully.")

	respBody, err := io.ReadAll(resp.Body)

	slog.Debug("Response body:", string(respBody))

	return string(respBody), err
}
