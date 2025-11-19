package scenarios

import (
	"log"
	"math/rand/v2"
	"time"

	"github.com/RafaelRochaS/edge-device-simulator/models"
	"github.com/RafaelRochaS/edge-device-simulator/utils"
	"gonum.org/v1/gonum/stat/distuv"
)

func ScenarioOne(config models.Config) {
	source := rand.NewPCG(uint64(config.BaseSeed), uint64(config.DeviceId))

	distExpo := distuv.Exponential{
		Rate: config.ArrivalRate,
		Src:  source,
	}
	distLogNormal := distuv.LogNormal{
		Mu:    float64(config.WorkloadMean),
		Sigma: float64(config.WorkloadStdVar),
		Src:   source,
	}

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
