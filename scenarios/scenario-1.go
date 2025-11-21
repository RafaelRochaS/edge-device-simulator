package scenarios

import (
	"log"

	"github.com/RafaelRochaS/edge-device-simulator/models"
	"github.com/RafaelRochaS/edge-device-simulator/utils"
)

func ScenarioOne(config models.Config) {
	scenarioWrapper(
		config,
		func(input ScenarioInput) {
			task := generateTask(input.config, input.DistLogNormal)

			log.Println("Creating task: ", task)

			err := utils.OffloadTask(config, *task)

			if err != nil {
				log.Println("Failed to offload task: ", err)
			}
		})
}
