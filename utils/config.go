package utils

import (
	"flag"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/RafaelRochaS/edge-device-simulator/models"
)

func GetConfig() (config models.Config) {
	scenarioFlag := flag.Int("scenario", 2, "Scenario to run:\n0 - Local processing\n1 - Cloud processing\n2 - Hybrid edge with xApp")
	lambdaFlag := flag.Float64("arrival-rate", 1.1, "Arrival rate of workloads in requests per second.")
	callbackFlag := flag.String("callback", "http://localhost:8080", "Callback URL to send results to.")
	durationFlag := flag.Duration("duration", time.Minute, "Time in seconds to run the simulation.")
	loadMeanFlag := flag.Int("workload-mean", 22, "Mean of workload sizes.")
	loadStdVarFlag := flag.Int("workload-std-var", 2, "Standard deviation of workload sizes.")
	kubeconfigPath := flag.String("kubeconfig", "./kubeconfig", "Path to kubeconfig file.")

	flag.Parse()

	if *scenarioFlag < 0 || *scenarioFlag > 2 {
		config.Scenario = models.MEC
	}
	config.Scenario = models.Scenario(*scenarioFlag)

	config.ArrivalRate = *lambdaFlag
	config.Callback = *callbackFlag
	config.Duration = *durationFlag
	config.WorkloadMean = *loadMeanFlag
	config.WorkloadStdVar = *loadStdVarFlag
	config.KubeconfigPath = *kubeconfigPath

	seed, err := strconv.Atoi(os.Getenv("BASE_SEED"))

	if err != nil {
		log.Fatal("Failed to parse BASE_SEED: ", err)
	}

	config.BaseSeed = seed

	deviceId, err := strconv.Atoi(os.Getenv("DEVICE_ID"))

	if err != nil {
		deviceId = -1
	}

	config.DeviceId = deviceId

	return
}
