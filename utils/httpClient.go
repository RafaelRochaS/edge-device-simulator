package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func makePostCall(data any, url string) (string, error) {
	body, err := json.Marshal(data)

	if err != nil {
		log.Fatal("Failed to parse body: ", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))

	if err != nil {
		log.Fatal("Failed to send callback: ", err)
	}

	log.Println("Response status:", resp.StatusCode)
	log.Println("Callback sent successfully.")

	respBody, err := io.ReadAll(resp.Body)

	log.Println("Response body:", string(respBody))

	return string(respBody), err
}
