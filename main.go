package main

import (
	"log"
	"os"
	"sort"
	"strconv"
	"time"

	"gonum.org/v1/gonum/stat/distuv"
)

func main() {
	id := os.Getenv("DEVICE_ID")

	parsedId, err := strconv.Atoi(id)
	if err != nil {
		parsedId = 0
	}

	log.Println("Starting edge device simulator with id:", parsedId)

	dist := distuv.Poisson{Lambda: 5}

	maxAttempts := 100
	workLoadSizes := make([]int, maxAttempts)
	durations := make([]time.Duration, maxAttempts)

	for i := 0; i < maxAttempts; i++ {
		n := int(dist.Rand()*50_000_000) + 1
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
