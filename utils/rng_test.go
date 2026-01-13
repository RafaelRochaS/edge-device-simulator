package utils

import (
	"testing"

	"github.com/RafaelRochaS/edge-device-simulator/models"
)

func TestGetDistributions(t *testing.T) {
	config := models.Config{
		BaseSeed:       42,
		DeviceId:       1,
		ArrivalRate:    0.8,
		WorkloadMean:   15,
		WorkloadStdVar: 2,
	}

	distExpo, distLogNormal := GetDistributions(config)

	// Verify they are not zero-initialized
	if distExpo.Rate != 0.8 {
		t.Errorf("expected rate 0.8, got %v", distExpo.Rate)
	}
	if distLogNormal.Mu != 15 {
		t.Errorf("expected Mu 15, got %v", distLogNormal.Mu)
	}

	// Verify reproducibility
	v1 := distExpo.Rand()
	v2 := distLogNormal.Rand()

	distExpo2, distLogNormal2 := GetDistributions(config)
	if distExpo2.Rand() != v1 {
		t.Errorf("Exponential distribution not reproducible")
	}
	if distLogNormal2.Rand() != v2 {
		t.Errorf("LogNormal distribution not reproducible")
	}
}
