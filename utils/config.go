package utils

import (
	"flag"
	"log/slog"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/RafaelRochaS/edge-device-simulator/models"
)

func GetConfig() (config models.Config) {
	slog.Debug("Parsing command line arguments...")

	scenario := flag.Int("scenario", 2, "Scenario to run:\n0 - Local processing\n1 - Cloud processing\n2 - Hybrid edge with xApp")
	lambda := flag.Float64("arrival-rate", 0.8, "Arrival rate of workloads in requests per second.")
	callback := flag.String("callback", "http://localhost:8080", "Callback URL to send results to.")
	duration := flag.Duration("duration", time.Minute, "Duration of the simulation.")
	loadMean := flag.Int("workload-mean", 15, "Mean of workload sizes.")
	loadStdVar := flag.Int("workload-std-var", 2, "Standard deviation of workload sizes.")
	kubeconfigPath := flag.String("kubeconfig", "./kubeconfig", "Path to kubeconfig file.")
	taskImage := flag.String("task-image", "task-sim", "Task image to run on Kubernetes. Not used on scenario 0.")
	taskImageRepository := flag.String("task-image-repo", "rafaelrs94/xapp-mec", "Docker repository to pull task image from. Not used on scenario 0.")
	k8sOffloadNamespace := flag.String("k8s-offload-ns", "task-offload", "Namespace to offload tasks to. Not used on scenario 0.")
	offloadThreshold := flag.Int("offload-threshold", 100_000, "Max task size to execute locally. Tasks greater are offloaded to the MEC handler. Used only on scenario 2.")
	mecHandlerAddr := flag.String("mec-handler-addr", "http://mec-handler:8080", "MEC handler address. Used only on scenario 2.")
	logLevel := flag.String("log-level", "info", "Log level. Valid values: debug, info, warn, error.")

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
	config.LogLevel = parseLogLevel(*logLevel)

	slog.SetLogLoggerLevel(config.LogLevel)

	seed, err := strconv.Atoi(os.Getenv("BASE_SEED"))

	if err != nil {
		slog.Error("Failed to parse BASE_SEED: ", err)
	}

	config.BaseSeed = seed

	deviceId, err := strconv.Atoi(os.Getenv("DEVICE_ID"))

	if err != nil {
		deviceId = -1
	}

	config.DeviceId = deviceId

	return
}

func parseLogLevel(flagVal string) slog.Level {
	switch strings.ToLower(flagVal) {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
