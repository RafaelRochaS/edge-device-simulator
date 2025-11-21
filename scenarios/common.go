package scenarios

import (
	"fmt"
	"log"
	"os"
	"sync"
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
	distExpo, _ := utils.GetDistributions(config)
	timeout := time.After(config.Duration)

	var wg sync.WaitGroup

execution:
	for {
		select {
		case <-timeout:
			break execution
		default:

			wg.Go(func() {
				runnerDistExpo, runnerDistLogNormal := utils.GetDistributions(config)
				runner(ScenarioInput{
					config, runnerDistExpo, runnerDistLogNormal,
				})
			})

			sleepTime := time.Duration(distExpo.Rand() * float64(time.Second))

			log.Printf("Task running, sleeping for %v ...\n", sleepTime)
			time.Sleep(sleepTime)
		}
	}

	log.Println("Finished execution, waiting for running jobs.")
	wg.Wait()
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
