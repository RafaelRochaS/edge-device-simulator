package utils

import (
	"context"
	"fmt"
	"log"

	"github.com/RafaelRochaS/edge-device-simulator/models"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func getClusterClientSetConfig(path string) (*kubernetes.Clientset, error) {
	config, err := clientcmd.BuildConfigFromFlags("", path)

	if err != nil {
		return nil, err
	}

	return kubernetes.NewForConfig(config)
}

func OffloadTask(config models.Config, task models.Task) error {
	client, err := getClusterClientSetConfig(config.KubeconfigPath)

	if err != nil {
		return err
	}

	pod := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: task.Id,
			Labels: map[string]string{
				"offload":  "true",
				"deviceId": task.DeviceId,
			},
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:  fmt.Sprintf("Offload: %s-%s", task.DeviceId, task.Id),
					Image: fmt.Sprintf("%s/%s", "localrepository:5000", task.Image),
					Env: []v1.EnvVar{{
						Name:  "WORKLOAD_SIZE",
						Value: "200",
					}, {
						Name:  "DEVICE_ID",
						Value: task.DeviceId}},
				},
			},
		},
	}

	log.Println("Offloading task: ", task.Id)

	_, err = client.CoreV1().Pods("offload-task").Create(context.TODO(), pod, metav1.CreateOptions{})

	return err
}
