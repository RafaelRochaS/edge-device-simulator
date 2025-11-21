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
	scenario := flag.Int("scenario", 2, "Scenario to run:\n0 - Local processing\n1 - Cloud processing\n2 - Hybrid edge with xApp")
	lambda := flag.Float64("arrival-rate", 1.1, "Arrival rate of workloads in requests per second.")
	callback := flag.String("callback", "http://localhost:8080", "Callback URL to send results to.")
	duration := flag.Duration("duration", time.Minute, "Time in seconds to run the simulation.")
	loadMean := flag.Int("workload-mean", 18, "Mean of workload sizes.")
	loadStdVar := flag.Int("workload-std-var", 2, "Standard deviation of workload sizes.")
	kubeconfigPath := flag.String("kubeconfig", "./kubeconfig", "Path to kubeconfig file.")
	taskImage := flag.String("task-image", "task-sim", "Task image to run on Kubernetes. Not used on scenario 0.")
	taskImageRepository := flag.String("task-image-repo", "rafaelrs94/xapp-mec", "Docker repository to pull task image from. Not used on scenario 0.")
	k8sOffloadNamespace := flag.String("k8s-offload-ns", "task-offload", "Namespace to offload tasks to. Not used on scenario 0.")
	offloadThreshold := flag.Int("offload-threshold", 100_000, "Max task size to execute locally. Tasks greater are offloaded to the MEC handler. Used only on scenario 2.")
	mecHandlerAddr := flag.String("mec-handler-addr", "http://mec-handler:8080", "MEC handler address. Used only on scenario 2.")

	flag.Parse()

	if *scenario < 0 || *scenario > 2 {
		config.Scenario = models.MEC
	}
	config.Scenario = models.Scenario(*scenario)

	config.ArrivalRate = *lambda
	config.Callback = *callback
	config.Duration = *duration
	config.WorkloadMean = *loadMean
	config.WorkloadStdVar = *loadStdVar
	config.KubeconfigPath = *kubeconfigPath
	config.TaskImage = *taskImage
	config.TaskImageRepository = *taskImageRepository
	config.K8sOffloadNamespace = *k8sOffloadNamespace
	config.MECOffloadThreshold = *offloadThreshold
	config.MECHandlerAddr = *mecHandlerAddr

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
