package scenarios

import (
	"testing"
	"time"

	"github.com/RafaelRochaS/edge-device-simulator/models"
)

func TestScenarioZero(t *testing.T) {
	config := models.Config{
		Duration:       50 * time.Millisecond,
		ArrivalRate:    100,
		WorkloadMean:   1,
		WorkloadStdVar: 1,
		BaseSeed:       42,
		DeviceId:       1,
		LocalCallback:  "http://localhost:8080",
	}

	// Just verifying no panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("ScenarioZero panicked: %v", r)
		}
	}()

	ScenarioZero(config)
}

func TestScenarioOne(t *testing.T) {
	config := models.Config{
		Duration:            50 * time.Millisecond,
		ArrivalRate:         100,
		WorkloadMean:        1,
		WorkloadStdVar:      1,
		BaseSeed:            42,
		DeviceId:            1,
		KubeconfigPath:      "non-existent-path",
		K8sOffloadNamespace: "default",
	}

	// Just verifying no panic.
	// OffloadTask will fail because of non-existent kubeconfig, but it should log and return error, not panic.
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("ScenarioOne panicked: %v", r)
		}
	}()

	ScenarioOne(config)
}

func TestScenarioTwo(t *testing.T) {
	config := models.Config{
		Duration:            50 * time.Millisecond,
		ArrivalRate:         100,
		WorkloadMean:        100,
		WorkloadStdVar:      1,
		BaseSeed:            42,
		DeviceId:            1,
		MECOffloadThreshold: 50,
		MECHandlerAddr:      "http://localhost:9999",
		LocalCallback:       "http://localhost:8080",
	}

	t.Run("Offload to MEC", func(t *testing.T) {
		// Set mean high to trigger offload
		c := config
		c.WorkloadMean = 100
		c.WorkloadStdVar = 1
		c.MECOffloadThreshold = 10

		defer func() {
			if r := recover(); r != nil {
				t.Errorf("ScenarioTwo panicked during MEC offload: %v", r)
			}
		}()

		ScenarioTwo(c)
	})

	t.Run("Execute Locally", func(t *testing.T) {
		// Set threshold high to trigger local execution
		c := config
		c.WorkloadMean = 1
		c.MECOffloadThreshold = 1000

		defer func() {
			if r := recover(); r != nil {
				t.Errorf("ScenarioTwo panicked during local execution: %v", r)
			}
		}()

		ScenarioTwo(c)
	})
}
