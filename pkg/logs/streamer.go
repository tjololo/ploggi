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
	return s.getLogReader(ctx, namespace, podname, containerName, true)
}

func (s Streamer) PodLogs(containerName, podname, namespace string, ctx context.Context) (string, error) {
	podLogs, err := s.getLogReader(ctx, namespace, podname, containerName, false)
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

func (s Streamer) getLogReader(ctx context.Context, namespace, podname, containername string, follow bool) (io.ReadCloser, error) {
	tailCount := int64(100)
	podLogOptions := corev1.PodLogOptions{
		Container: containername,
		Follow:    follow,
		TailLines: &tailCount,
	}

	podLogRrequest := s.Clientset.CoreV1().Pods(namespace).GetLogs(podname, &podLogOptions)
	return podLogRrequest.Stream(ctx)
}