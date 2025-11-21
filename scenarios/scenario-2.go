package scenarios

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/RafaelRochaS/edge-device-simulator/models"
	"github.com/RafaelRochaS/edge-device-simulator/utils"
	"github.com/google/uuid"
)

func ScenarioTwo(config models.Config) {
	distExpo, distLogNormal := utils.GetDistributions(config)
	timeout := time.After(config.Duration)

execution:
	for {
		select {
		case <-timeout:
			break execution
		default:
			task := new(models.Task)

			task.Workload = int(distLogNormal.Rand())
			task.DeviceId = os.Getenv("DEVICE_ID")
			task.Image = fmt.Sprintf("%s:%s", config.TaskImageRepository, config.TaskImage)
			task.CPU = "1"
			task.Mem = 512
			task.CallbackUrl = config.Callback
			task.Id = uuid.New().String()

			log.Println("Creating task: ", task)

			if task.Workload > config.MECOffloadThreshold {
				err := utils.OffloadTask(config, *task)

				if err != nil {
					log.Println("Failed to offload task: ", err)
				}

				continue
			}

			go func() {
				callbackData := utils.GetCallbackData()
				callbackData.DeviceID = config.DeviceId

				duration := utils.CpuBoundWork(task.Workload)
				callbackData.Duration = duration.Seconds()

				utils.SendCallback(callbackData, config.Callback)
			}()

			sleepTime := distExpo.Rand() * time.Second.Seconds()

			log.Printf("Task offloaded, sleeping for %f seconds...\n", sleepTime)
			time.Sleep(time.Duration(sleepTime))
		}
	}
}
