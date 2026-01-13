package utils

import (
	"log/slog"
	"os"
	"testing"

	"github.com/RafaelRochaS/edge-device-simulator/models"
)

func TestParseLogLevel(t *testing.T) {
	tests := []struct {
		input    string
		expected slog.Level
	}{
		{"debug", slog.LevelDebug},
		{"DEBUG", slog.LevelDebug},
		{"info", slog.LevelInfo},
		{"warn", slog.LevelWarn},
		{"error", slog.LevelError},
		{"invalid", slog.LevelInfo},
	}

	for _, tt := range tests {
		result := parseLogLevel(tt.input)
		if result != tt.expected {
			t.Errorf("parseLogLevel(%s) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestGetConfig(t *testing.T) {
	// Backup original args and env
	origArgs := os.Args
	defer func() { os.Args = origArgs }()

	origBaseSeed := os.Getenv("BASE_SEED")
	origDeviceId := os.Getenv("DEVICE_ID")
	defer func() {
		os.Setenv("BASE_SEED", origBaseSeed)
		os.Setenv("DEVICE_ID", origDeviceId)
	}()

	os.Setenv("BASE_SEED", "42")
	os.Setenv("DEVICE_ID", "1")

	// Set command line flags
	os.Args = []string{
		"cmd",
		"-scenario", "1",
		"-arrival-rate", "0.5",
		"-workload-mean", "20",
		"-log-level", "debug",
	}

	// Reset flags to allow re-parsing (flag.Parse is called inside GetConfig)
	// Since flag.Parse is called inside GetConfig, and we can't easily reset the global FlagSet
	// without reaching into internals, we should be careful.
	// However, GetConfig calls flag.Parse(), so we must ensure it's not already parsed
	// or use a separate FlagSet in the production code if we wanted it to be more testable.
	// For now, let's see if it works once.

	config := GetConfig()

	if config.Scenario != models.Cloud {
		t.Errorf("expected scenario %v, got %v", models.Cloud, config.Scenario)
	}
	if config.ArrivalRate != 0.5 {
		t.Errorf("expected arrival rate 0.5, got %v", config.ArrivalRate)
	}
	if config.WorkloadMean != 20 {
		t.Errorf("expected workload mean 20, got %v", config.WorkloadMean)
	}
	if config.BaseSeed != 42 {
		t.Errorf("expected base seed 42, got %v", config.BaseSeed)
	}
	if config.DeviceId != 1 {
		t.Errorf("expected device id 1, got %v", config.DeviceId)
	}
	if config.LogLevel != slog.LevelDebug {
		t.Errorf("expected log level debug, got %v", config.LogLevel)
	}
}
