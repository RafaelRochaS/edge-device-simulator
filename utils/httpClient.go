package utils

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func makePostCall(data any, url string) error {
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

	return nil
}
