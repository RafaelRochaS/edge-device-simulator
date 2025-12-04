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

func (c Config) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("scenario", c.Scenario.String()),
		slog.String("callback", c.Callback),
		slog.Float64("arrivalRate", c.ArrivalRate),
		slog.Duration("duration", c.Duration),
		slog.Int("workloadMean", c.WorkloadMean),
		slog.Int("workloadStdVar", c.WorkloadStdVar),
		slog.Int("baseSeed", c.BaseSeed),
		slog.Int("deviceId", c.DeviceId),
		slog.String("kubeconfigPath", c.KubeconfigPath),
		slog.String("taskImage", c.TaskImage),
		slog.String("taskImageRepository", c.TaskImageRepository),
		slog.String("k8sOffloadNamespace", c.K8sOffloadNamespace),
		slog.Int("mecOffloadThreshold", c.MECOffloadThreshold),
		slog.String("mecHandlerAddr", c.MECHandlerAddr),
		slog.Int("logLevel", int(c.LogLevel)),
	)
}
