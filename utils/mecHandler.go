package utils

import (
	"log"

	"github.com/RafaelRochaS/edge-device-simulator/models"
)

func MECOffload(task models.Task, url string) error {
	log.Println("Offloading to MEC handler task: ", task)
	return makePostCall(task, url)
}
