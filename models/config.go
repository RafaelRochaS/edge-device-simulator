package models

import (
	"log/slog"
	"time"
)

type Config struct {
	Scenario            Scenario
	Callback            string
	ArrivalRate         float64
	Duration            time.Duration
	WorkloadMean        int
	WorkloadStdVar      int
	BaseSeed            int
	DeviceId            int
	KubeconfigPath      string
	TaskImage           string
	TaskImageRepository string
	K8sOffloadNamespace string
	MECOffloadThreshold int
	MECHandlerAddr      string
	LogLevel            slog.Level
}
