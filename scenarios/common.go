package scenarios

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
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
	slog.Debug("Running scenario wrapper")

	distExpo, _ := utils.GetDistributions(config)
	timeout := time.After(config.Duration)

	var wg sync.WaitGroup

	slog.Debug("Starting execution loop...")

execution:
	for {
		select {
		case <-timeout:
			break execution
		default:

			wg.Go(func() {
				slog.Debug("Starting task execution...")
				runnerDistExpo, runnerDistLogNormal := utils.GetDistributions(config)
				runner(ScenarioInput{
					config, runnerDistExpo, runnerDistLogNormal,
				})
			})

			sleepTime := time.Duration(distExpo.Rand() * float64(time.Second))

			slog.Info("Task running, sleeping for ", slog.Duration("sleepTime", sleepTime))
			time.Sleep(sleepTime)
		}
	}

	slog.Info("Finished execution, waiting for running jobs.")
	wg.Wait()
}

func generateTask(config models.Config, distLogNormal distuv.LogNormal) *models.Task {
	slog.Debug("Generating task...")
	task := new(models.Task)

	deviceId, err := strconv.Atoi(os.Getenv("DEVICE_ID"))

	if err != nil {
		deviceId = -1
	}

	task.Workload = int(distLogNormal.Rand())
	task.DeviceId = deviceId
	task.Image = fmt.Sprintf("%s:%s", config.TaskImageRepository, config.TaskImage)
	task.CPU = "1"
	task.Mem = 512
	task.CallbackUrl = config.Callback
	task.Id = uuid.New().String()

	slog.Debug("Generated task: ", slog.Any("task", task))

	return task
}

func executeTask(config models.Config, distLogNormal distuv.LogNormal) {
	slog.Debug("Executing task...")

	callbackData := utils.GetCallbackData()
	callbackData.DeviceID = config.DeviceId

	n := int(distLogNormal.Rand())
	callbackData.WorkloadSize = n

	slog.Debug("Starting workload of size", slog.Int("workload", n))

	duration := utils.CpuBoundWork(n)
	callbackData.Duration = duration.Seconds()

	slog.Debug("Generated callback data: ", slog.Any("callbackData", callbackData))

	utils.SendCallback(callbackData, config.Callback)
}
