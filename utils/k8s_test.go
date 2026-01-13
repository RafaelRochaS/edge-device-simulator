package utils

import (
	"testing"

	"github.com/RafaelRochaS/edge-device-simulator/models"
)

func TestOffloadTask_InvalidConfig(t *testing.T) {
	config := models.Config{
		KubeconfigPath: "non-existent-path",
	}
	task := models.Task{
		Id: "test-task",
	}

	err := OffloadTask(config, task)
	if err == nil {
		t.Error("expected error due to missing kubeconfig")
	}
}

func TestGetClusterClientSetConfig_InvalidPath(t *testing.T) {
	_, err := getClusterClientSetConfig("non-existent-path")
	if err == nil {
		t.Error("expected error for invalid kubeconfig path")
	}
}
