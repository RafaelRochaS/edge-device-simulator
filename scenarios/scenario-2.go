package scenarios

import (
	"log"

	"github.com/RafaelRochaS/edge-device-simulator/models"
	"github.com/RafaelRochaS/edge-device-simulator/utils"
)

func ScenarioTwo(config models.Config) {
	scenarioWrapper(
		config,
		func(input ScenarioInput) {
			task := generateTask(input.config, input.DistLogNormal)
			log.Println("Creating task: ", task)

			if task.Workload > config.MECOffloadThreshold {
				err := utils.MECOffload(*task, config.MECHandlerAddr)

				if err != nil {
					log.Println("Failed to offload task: ", err)
				}

			} else {
				go executeTask(config, input.DistLogNormal)
			}
		})
}
