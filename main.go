package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"log"
	"math/rand/v2"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
	"gonum.org/v1/gonum/stat/distuv"
)

type Scenario int

const (
	Local Scenario = iota
	Cloud
	MEC
)

type Config struct {
	Scenario       Scenario
	Callback       string
	ArrivalRate    float64
	Duration       time.Duration
	WorkloadMean   int
	WorkloadStdVar int
	BaseSeed       int
	DeviceId       int
}

type CallbackData struct {
	TaskID        string  `json:"taskId"`
	DeviceID      int     `json:"deviceId"`
	WorkloadSize  int     `json:"workloadSize"`
	ExecutionSite string  `json:"executionSite"`
	CreatedAt     int64   `json:"createdAt"`
	Duration      float64 `json:"duration"`
}

func main() {
	log.Println("Starting edge device simulator...")

	config := getConfig()
	log.Printf("Loaded config: %+v\n", config)

	log.Println("Starting scenario...")

	switch config.Scenario {
	case Local:
		scenarioOne(config)
	default:
		log.Fatalf("Scenario %d not yet implemented.", config.Scenario)
	}

	log.Println("Finished scenario.")
}

func getConfig() (config Config) {
	scenarioFlag := flag.Int("scenario", 2, "Scenario to run:\n0 - Local processing\n1 - Cloud processing\n2 - Hybrid edge with xApp")
	lambdaFlag := flag.Float64("arrival-rate", 2, "Arrival rate of workloads in requests per second.")
	callbackFlag := flag.String("callback", "http://localhost:8080", "Callback URL to send results to.")
	durationFlag := flag.Duration("duration", time.Minute, "Time in seconds to run the simulation.")
	loadMeanFlag := flag.Int("workload-mean", 25, "Mean of workload sizes.")
	loadStdVarFlag := flag.Int("workload-std-var", 3, "Standard deviation of workload sizes.")

	flag.Parse()

	if *scenarioFlag < 0 || *scenarioFlag > 2 {
		config.Scenario = MEC
	}
	config.Scenario = Scenario(*scenarioFlag)

	config.ArrivalRate = *lambdaFlag
	config.Callback = *callbackFlag
	config.Duration = *durationFlag
	config.WorkloadMean = *loadMeanFlag
	config.WorkloadStdVar = *loadStdVarFlag

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

func scenarioOne(config Config) {
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
			callbackData := getCallbackData()
			callbackData.DeviceID = config.DeviceId

			n := int(distLogNormal.Rand())
			callbackData.WorkloadSize = n

			log.Printf("Starting workload of size %d...\n", n)

			duration := CpuBoundWork(n)
			callbackData.Duration = duration.Seconds()

			sendCallback(callbackData, config.Callback)

			sleepTime := distExpo.Rand() * time.Second.Seconds()

			log.Printf("Finished, sleeping for %f seconds...\n", sleepTime)
			time.Sleep(time.Duration(sleepTime))
		}
	}
}

func getCallbackData() CallbackData {
	return CallbackData{
		TaskID:        uuid.New().String(),
		ExecutionSite: "local",
		CreatedAt:     time.Now().Unix(),
	}
}

func sendCallback(data CallbackData, url string) {
	log.Printf("Sending callback: %+v\n", data)

	body, err := json.Marshal(data)

	if err != nil {
		log.Fatal("Failed to parse body: ", err)
	}

	log.Println("Body:", string(body))

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))

	if err != nil {
		log.Fatal("Failed to send callback: ", err)
	}

	log.Println("Response status:", resp.StatusCode)
	log.Println("Callback sent successfully.")
}
