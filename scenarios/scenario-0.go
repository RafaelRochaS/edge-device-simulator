package scenarios

import (
	"log"
	"time"

	"github.com/RafaelRochaS/edge-device-simulator/models"
	"github.com/RafaelRochaS/edge-device-simulator/utils"
)

func ScenarioZero(config models.Config) {
	distExpo, distLogNormal := utils.GetDistributions(config)
	timeout := time.After(config.Duration)

execution:
	for {
		select {
		case <-timeout:
			break execution
		default:
			callbackData := utils.GetCallbackData()
			callbackData.DeviceID = config.DeviceId

			n := int(distLogNormal.Rand())
			callbackData.WorkloadSize = n

			log.Printf("Starting workload of size %d...\n", n)

			duration := utils.CpuBoundWork(n)
			callbackData.Duration = duration.Seconds()

			utils.SendCallback(callbackData, config.Callback)

			sleepTime := distExpo.Rand() * time.Second.Seconds()

			log.Printf("Finished, sleeping for %f seconds...\n", sleepTime)
			time.Sleep(time.Duration(sleepTime))
		}
	}
}
