package main

import (
	"log/slog"

	"github.com/RafaelRochaS/edge-device-simulator/models"
	"github.com/RafaelRochaS/edge-device-simulator/scenarios"
	"github.com/RafaelRochaS/edge-device-simulator/utils"
)

func main() {
	config := utils.GetConfig()

	slog.Info("Starting edge device simulator...")
	slog.Debug("Loaded config: %+v\n", config)
	slog.Info("Starting scenario...")

	switch config.Scenario {
	case models.Local:
		scenarios.ScenarioZero(config)
	case models.Cloud:
		scenarios.ScenarioOne(config)
	case models.MEC:
		scenarios.ScenarioTwo(config)
	default:
		slog.Error("Scenario %d not yet implemented.", config.Scenario)
	}

	slog.Info("Finished simulation.")
}
