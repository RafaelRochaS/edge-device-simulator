package scenarios

import (
	"os"
	"sync/atomic"
	"testing"
	"time"

	"github.com/RafaelRochaS/edge-device-simulator/models"
	"gonum.org/v1/gonum/stat/distuv"
)

func TestGenerateTask(t *testing.T) {
	config := models.Config{
		TaskImageRepository: "repo",
		TaskImage:           "image",
		Callback:            "http://callback",
	}
	distLogNormal := distuv.LogNormal{Mu: 1, Sigma: 0.1}

	t.Run("with DEVICE_ID", func(t *testing.T) {
		err := os.Setenv("DEVICE_ID", "123")
		if err != nil {
			return
		}
		defer func() {
			err := os.Unsetenv("DEVICE_ID")
			if err != nil {
				t.Error(err)
			}
		}()

		task := generateTask(config, distLogNormal)

		if task.DeviceId != 123 {
			t.Errorf("expected DeviceId 123, got %d", task.DeviceId)
		}
		if task.Image != "repo:image" {
			t.Errorf("expected Image repo:image, got %s", task.Image)
		}
		if task.CallbackUrl != config.Callback {
			t.Errorf("expected CallbackUrl %s, got %s", config.Callback, task.CallbackUrl)
		}
		if task.Id == "" {
			t.Error("expected Task Id to be generated, got empty string")
		}
	})

	t.Run("without DEVICE_ID", func(t *testing.T) {
		err := os.Unsetenv("DEVICE_ID")
		if err != nil {
			return
		}

		task := generateTask(config, distLogNormal)

		if task.DeviceId != -1 {
			t.Errorf("expected DeviceId -1 when DEVICE_ID is unset, got %d", task.DeviceId)
		}
	})

	t.Run("invalid DEVICE_ID", func(t *testing.T) {
		err := os.Setenv("DEVICE_ID", "abc")
		if err != nil {
			return
		}
		defer func() {
			err := os.Unsetenv("DEVICE_ID")
			if err != nil {

			}
		}()

		task := generateTask(config, distLogNormal)

		if task.DeviceId != -1 {
			t.Errorf("expected DeviceId -1 when DEVICE_ID is invalid, got %d", task.DeviceId)
		}
	})
}

func TestExecuteTask(t *testing.T) {
	// executeTask calls utils.CpuBoundWork and utils.SendCallback.
	// We can't easily mock these without changing the code,
	// but we can at least ensure it runs with valid inputs.
	// Note: utils.SendCallback will try to make a network call,
	// but it logs a warning if it fails, it doesn't panic.

	config := models.Config{
		DeviceId:      1,
		LocalCallback: "http://localhost:8080",
	}
	distLogNormal := distuv.LogNormal{Mu: 1, Sigma: 0.1}

	// Just verifying no panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("executeTask panicked: %v", r)
		}
	}()

	executeTask(config, distLogNormal)
}

func TestScenarioWrapper(t *testing.T) {
	config := models.Config{
		Duration:       50 * time.Millisecond,
		ArrivalRate:    100, // High rate for fast tasks
		WorkloadMean:   1,
		WorkloadStdVar: 1,
		BaseSeed:       42,
		DeviceId:       1,
	}

	var callCount int32
	runner := func(input ScenarioInput) {
		atomic.AddInt32(&callCount, 1)
	}

	scenarioWrapper(config, runner)

	if atomic.LoadInt32(&callCount) == 0 {
		t.Error("expected runner to be called at least once")
	}
}
