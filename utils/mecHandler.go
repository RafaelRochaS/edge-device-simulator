package utils

import (
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/RafaelRochaS/edge-device-simulator/models"
)

const TasksEndpoint = "/ric/v1/mec/tasks"
const OffloadEndpoint = "/ric/v1/mec/start"

func MECOffload(task models.Task, url string) error {
	slog.Info("Offloading to MEC handler task: ", slog.Any("task", task))

	offloadTaskId, err := registerTask(task, url)

	if err != nil {
		slog.Error("Failed to offload task: ", slog.Any("error", err))
		return err
	}

	return startTask(task.DeviceId, offloadTaskId, url)
}

func registerTask(task models.Task, url string) (string, error) {
	slog.Debug("Registering task: ", slog.Any("task", task))
	createdTask, err := makePostCall(task, fmt.Sprintf("%s%s", url, TasksEndpoint))

	if err != nil {
		slog.Error("Failed to register task: ", slog.Any("error", err))
		return "", err
	}

	var taskResponse models.RegisterTaskResponse

	err = json.Unmarshal([]byte(createdTask), &taskResponse)

	slog.Debug("Task response: ", slog.Any("taskResponse", taskResponse))

	if err != nil {
		slog.Error("Failed to parse task response: ", slog.Any("error", err))
		return "", err
	}

	slog.Debug("Task registered successfully: ", slog.Any("taskResponse", taskResponse))
	return taskResponse.Id, nil
}

func startTask(deviceId int, taskId, url string) error {
	slog.Info("Starting task: ", slog.String("taskId", taskId))

	startTaskRequest := models.StartTaskRequest{
		Id:       taskId,
		DeviceId: deviceId,
	}

	slog.Debug("Starting task request: ", slog.Any("taskRequest", startTaskRequest))

	_, err := makePostCall(startTaskRequest, fmt.Sprintf("%s%s", url, OffloadEndpoint))

	if err != nil {
		slog.Error("Failed to start task: ", slog.Any("error", err))
	}

	return err
}
