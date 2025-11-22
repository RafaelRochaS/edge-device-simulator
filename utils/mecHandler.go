package utils

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/RafaelRochaS/edge-device-simulator/models"
)

const TasksEndpoint = "/ric/v1/mec/tasks"
const OffloadEndpoint = "/ric/v1/mec/start"

func MECOffload(task models.Task, url string) error {
	log.Println("Offloading to MEC handler task: ", task)

	offloadTaskId, err := registerTask(task, url)

	if err != nil {
		return err
	}

	return startTask(task.DeviceId, offloadTaskId, url)
}

func registerTask(task models.Task, url string) (string, error) {
	log.Println("Registering task: ", task)
	createdTask, err := makePostCall(task, fmt.Sprintf("%s%s", url, TasksEndpoint))

	if err != nil {
		return "", err
	}

	var taskResponse models.RegisterTaskResponse

	err = json.Unmarshal([]byte(createdTask), &taskResponse)

	if err != nil {
		return "", err
	}

	log.Println("Task registered successfully: ", taskResponse)
	return taskResponse.Id, nil
}

func startTask(deviceId int, taskId, url string) error {
	log.Println("Starting task: ", taskId)

	startTaskRequest := models.StartTaskRequest{
		Id:       taskId,
		DeviceId: deviceId,
	}

	_, err := makePostCall(startTaskRequest, fmt.Sprintf("%s%s", url, OffloadEndpoint))

	return err
}
