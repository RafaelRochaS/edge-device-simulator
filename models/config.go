package models

import "time"

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
}
