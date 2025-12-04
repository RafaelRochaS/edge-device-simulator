package main

import (
	"log/slog"
	"os"

	"github.com/RafaelRochaS/edge-device-simulator/models"
	"github.com/RafaelRochaS/edge-device-simulator/scenarios"
	"github.com/RafaelRochaS/edge-device-simulator/utils"
)

func main() {
	config := utils.GetConfig()

	slog.Info("Starting edge device simulator...")
	slog.Debug("Loaded:\n", "config", config)
	slog.Info("Starting scenario...")

	switch config.Scenario {
	case models.Local:
		scenarios.ScenarioZero(config)
	case models.Cloud:
		scenarios.ScenarioOne(config)
	case models.MEC:
		scenarios.ScenarioTwo(config)
	default:
		slog.Error("Scenario %d not yet implemented.", slog.Int("scenario", int(config.Scenario)))
		os.Exit(1)
	}

	slog.Info("Finished simulation.")
}
