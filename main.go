package main

import (
	"log"

	"github.com/RafaelRochaS/edge-device-simulator/models"
	"github.com/RafaelRochaS/edge-device-simulator/scenarios"
	"github.com/RafaelRochaS/edge-device-simulator/utils"
)

func main() {
	config := utils.GetConfig()

	log.Println("Starting edge device simulator...")
	log.Printf("Loaded config: %+v\n", config)
	log.Println("Starting scenario...")

	switch config.Scenario {
	case models.Local:
		scenarios.ScenarioOne(config)
	default:
		log.Fatalf("Scenario %d not yet implemented.", config.Scenario)
	}

	log.Println("Finished scenario.")
}
