package logs

import (
	"bytes"
	"context"
	"io"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
)

type Streamer struct {
	Clientset *kubernetes.Clientset
}

func (s Streamer) StreamPodLogs(containerName, podname, namespace string, ctx context.Context) (io.ReadCloser, error) {
	tailCount := int64(100)
	podLogOptions := corev1.PodLogOptions{
		Container: containerName,
		Follow:    true,
		TailLines: &tailCount,
	}

	podLogRrequest := s.Clientset.CoreV1().Pods(namespace).GetLogs(podname, &podLogOptions)
	return podLogRrequest.Stream(ctx)
}

func (s Streamer) PodLogs(containerName, podname, namespace string, ctx context.Context) (string, error) {
	tailCount := int64(1000)
	podLogOptions := corev1.PodLogOptions{
		Container: containerName,
		Follow:    false,
		TailLines: &tailCount,
	}

	podLogRrequest := s.Clientset.CoreV1().Pods(namespace).GetLogs(podname, &podLogOptions)
	podLogs, err := podLogRrequest.Stream(ctx)
	if err != nil {
        return "", err
    }
    defer podLogs.Close()

    buf := new(bytes.Buffer)
    _, err = io.Copy(buf, podLogs)
    if err != nil {
        return "", err
    }
    str := buf.String()

    return str, nil
}