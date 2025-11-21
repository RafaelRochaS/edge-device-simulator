package scenarios

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/RafaelRochaS/edge-device-simulator/models"
	"github.com/RafaelRochaS/edge-device-simulator/utils"
	"github.com/google/uuid"
	"gonum.org/v1/gonum/stat/distuv"
)

type ScenarioInput struct {
	config        models.Config
	DistExpo      distuv.Exponential
	DistLogNormal distuv.LogNormal
}

func scenarioWrapper(config models.Config, runner func(input ScenarioInput)) {
	distExpo, distLogNormal := utils.GetDistributions(config)
	timeout := time.After(config.Duration)

execution:
	for {
		select {
		case <-timeout:
			break execution
		default:

			runner(ScenarioInput{
				config, distExpo, distLogNormal,
			})

			sleepTime := distExpo.Rand() * time.Second.Seconds()

			log.Printf("Finished, sleeping for %f seconds...\n", sleepTime)
			time.Sleep(time.Duration(sleepTime))
		}
	}
}

func generateTask(config models.Config, distLogNormal distuv.LogNormal) *models.Task {
	task := new(models.Task)

	task.Workload = int(distLogNormal.Rand())
	task.DeviceId = os.Getenv("DEVICE_ID")
	task.Image = fmt.Sprintf("%s:%s", config.TaskImageRepository, config.TaskImage)
	task.CPU = "1"
	task.Mem = 512
	task.CallbackUrl = config.Callback
	task.Id = uuid.New().String()

	return task
}

func executeTask(config models.Config, distLogNormal distuv.LogNormal) {
	callbackData := utils.GetCallbackData()
	callbackData.DeviceID = config.DeviceId

	n := int(distLogNormal.Rand())
	callbackData.WorkloadSize = n

	log.Printf("Starting workload of size %d...\n", n)

	duration := utils.CpuBoundWork(n)
	callbackData.Duration = duration.Seconds()

	utils.SendCallback(callbackData, config.Callback)
}
