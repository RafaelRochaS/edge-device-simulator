package scenarios

import (
	"log/slog"

	"github.com/RafaelRochaS/edge-device-simulator/models"
	"github.com/RafaelRochaS/edge-device-simulator/utils"
)

func ScenarioOne(config models.Config) {
	scenarioWrapper(
		config,
		func(input ScenarioInput) {
			task := generateTask(input.config, input.DistLogNormal)

			slog.Info("Creating task: ", slog.Any("task", task))

			err := utils.OffloadTask(config, *task)

			if err != nil {
				slog.Error("Failed to offload task: ", slog.Any("error", err))
			}
		})
}
