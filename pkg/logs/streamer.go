package logs

import (
	"context"
	"io"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
)

type Streamer struct {
	clientset *kubernetes.Clientset
}

func (s Streamer) StreamPodLogs(containerName, podname, namespace string, ctx context.Context) (io.ReadCloser, error) {
	tailCount := int64(100)
	podLogOptions := corev1.PodLogOptions{
		Container: containerName,
		Follow:    true,
		TailLines: &tailCount,
	}

	podLogRrequest := s.clientset.CoreV1().Pods(namespace).GetLogs(podname, &podLogOptions)
	return podLogRrequest.Stream(ctx)
}
