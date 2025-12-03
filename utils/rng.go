package utils

import (
	"log/slog"
	"math/rand/v2"

	"github.com/RafaelRochaS/edge-device-simulator/models"
	"gonum.org/v1/gonum/stat/distuv"
)

func GetDistributions(config models.Config) (distuv.Exponential, distuv.LogNormal) {
	slog.Debug("Generating distributions...")
	source := rand.NewPCG(uint64(config.BaseSeed), uint64(config.DeviceId))

	distExpo := distuv.Exponential{
		Rate: config.ArrivalRate,
		Src:  source,
	}
	distLogNormal := distuv.LogNormal{
		Mu:    float64(config.WorkloadMean),
		Sigma: float64(config.WorkloadStdVar),
		Src:   source,
	}

	return distExpo, distLogNormal
}
