package main

import (
	"log"
	"math"
	"time"
)

func CpuBoundWork(n int) time.Duration {
	log.Println("Running work for value:", n)
	startTime := time.Now()

	acc := 0.0
	for i := 0; i < n; i++ {
		acc += math.Sqrt(float64(i)) * math.Sin(float64(i))
	}

	return time.Since(startTime)
}
