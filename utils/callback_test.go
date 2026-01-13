package utils

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RafaelRochaS/edge-device-simulator/models"
)

func TestGetCallbackData(t *testing.T) {
	data := GetCallbackData()
	if data.TaskID == "" {
		t.Error("expected non-empty TaskID")
	}
	if data.ExecutionSite != "local" {
		t.Errorf("expected ExecutionSite local, got %s", data.ExecutionSite)
	}
	if data.CreatedAt == 0 {
		t.Error("expected non-zero CreatedAt")
	}
}

func TestSendCallback(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST request, got %s", r.Method)
		}
		var data models.CallbackData
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			t.Errorf("failed to decode body: %v", err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	}))
	defer server.Close()

	data := models.CallbackData{TaskID: "test-task"}
	// This should not panic and should succeed
	SendCallback(data, server.URL)
}
