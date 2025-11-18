package main

import (
	"flag"
	"log"
	"sort"
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
	Scenario    Scenario
	Callback    string
	ArrivalRate float64
	Iterations  int
}

type CallbackData struct {
	TaskID        string  `json:"taskId"`
	DeviceID      string  `json:"deviceId"`
	WorkloadSize  int     `json:"workloadSize"`
	ExecutionSite string  `json:"executionSite"`
	CreatedAt     int64   `json:"createdAt"`
	Duration      float64 `json:"duration"`
}

func main() {
	id := uuid.New()
	log.Println("Starting edge device simulator with id:", id)

	config := getConfig()
	log.Printf("Loaded config: %+v\n", config)

	log.Println("Starting scenario...")

	switch config.Scenario {
	case Local:
		scenarioOne(config)
	default:
		log.Fatalf("Scenario %d not yet implemented.", config.Scenario)
	}
}

func getConfig() (config Config) {
	scenarioFlag := flag.Int("scenario", 3, "Scenario to run:\n1 - Local processing\n2 - Cloud processing\n3 - Hybrid edge with xApp")
	lambdaFlag := flag.Float64("arrival-rate", 0.5, "Arrival rate of workloads in requests per second.")
	callbackFlag := flag.String("callback", "http://localhost", "Callback URL to send results to.")
	iterationsFlag := flag.Int("duration", 100, "Total iterations of tasks to send.")

	flag.Parse()

	if *scenarioFlag < 1 || *scenarioFlag > 3 {
		config.Scenario = MEC
	}
	config.Scenario = Scenario(*scenarioFlag)

	config.ArrivalRate = *lambdaFlag
	config.Callback = *callbackFlag
	config.Iterations = *iterationsFlag

	return
}

func scenarioOne(config Config) {
	dist := distuv.Poisson{Lambda: config.ArrivalRate}

	workLoadSizes := make([]int, config.Iterations)
	durations := make([]time.Duration, config.Iterations)

	for i := 0; i < config.Iterations; i++ {
		n := int(dist.Rand()*50_000) + 1
		workLoadSizes[i] = n
		durations[i] = CpuBoundWork(n)
	}

	sort.Ints(workLoadSizes)
	sort.Slice(durations, func(i, j int) bool {
		return durations[i] < durations[j]
	})

	log.Println("Work load sizes:", workLoadSizes)
	log.Println("Work load durations:", durations)
}

func getCallbackData(workloadSize int) CallbackData {
	return CallbackData{
		TaskID:        uuid.New().String(),
		WorkloadSize:  workloadSize,
		ExecutionSite: "edge",
		CreatedAt:     time.Now().Unix(),
	}
}
