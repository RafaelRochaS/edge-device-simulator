package scenarios

import (
	"github.com/RafaelRochaS/edge-device-simulator/models"
)

func ScenarioZero(config models.Config) {
	scenarioWrapper(
		config,
		func(input ScenarioInput) {
			go executeTask(input.config, input.DistLogNormal)
		})
}
