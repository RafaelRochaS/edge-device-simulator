package scenarios

import (
	"log/slog"

	"github.com/RafaelRochaS/edge-device-simulator/models"
	"github.com/RafaelRochaS/edge-device-simulator/utils"
)

func ScenarioTwo(config models.Config) {
	scenarioWrapper(
		config,
		func(input ScenarioInput) {
			task := generateTask(input.config, input.DistLogNormal)
			slog.Info("Creating task: ", task)

			if task.Workload > config.MECOffloadThreshold {
				err := utils.MECOffload(*task, config.MECHandlerAddr)

				if err != nil {
					slog.Error("Failed to offload task: ", err)
				}

			} else {
				executeTask(config, input.DistLogNormal)
			}
		})
}
