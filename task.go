package main

import (
	"time"
)

func CpuBoundWork(n int) time.Duration {
	start := time.Now()

	x := uint64(0xDEADBEEF)
	for i := 0; i < n; i++ {
		x ^= uint64(i) * 0x9e3779b97f4a7c15
		x ^= x >> 33
		x *= 0xff51afd7ed558ccd
		x ^= x >> 33
		x *= 0xc4ceb9fe1a85ec53
		x ^= x >> 33
	}

	_ = x

	return time.Since(start)
}
