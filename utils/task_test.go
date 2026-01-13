package utils

import (
	"testing"
)

func TestCpuBoundWork(t *testing.T) {
	// Just verify it runs and returns a non-negative duration
	duration := CpuBoundWork(100)
	if duration < 0 {
		t.Errorf("expected non-negative duration, got %v", duration)
	}

	// Verify that more work takes (presumably) some time or at least doesn't crash
	duration2 := CpuBoundWork(1000)
	if duration2 < 0 {
		t.Errorf("expected non-negative duration, got %v", duration2)
	}
}
