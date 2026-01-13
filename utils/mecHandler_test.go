package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RafaelRochaS/edge-device-simulator/models"
)

func TestMECOffload(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case TasksEndpoint:
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(models.RegisterTaskResponse{Id: "mec-task-id"})
		case OffloadEndpoint:
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, "started")
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	task := models.Task{Id: "test-task", DeviceId: 1}
	err := MECOffload(task, server.URL)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestRegisterTask_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	task := models.Task{Id: "test-task"}
	_, err := registerTask(task, server.URL)

	if err == nil {
		t.Error("expected error from registerTask on 500 response")
	}
}

func TestStartTask_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	err := startTask(1, "task-id", server.URL)

	if err == nil {
		t.Error("expected error from startTask on 500 response")
	}
}
