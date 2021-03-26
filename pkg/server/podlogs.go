package server

import (
	"context"
	"fmt"
	"io"

	"github.com/tjololo/ploggi/pkg/api/ploggi"
	plogs "github.com/tjololo/ploggi/pkg/logs"
)

type PodLogServer struct {
	ploggi.UnimplementedPodLogsServer
	LogStreamer *plogs.Streamer
}

func (s *PodLogServer) GetLog(ctx context.Context, in *ploggi.Pod) (*ploggi.PodLog, error) {
	logs, err := s.LogStreamer.PodLogs(in.Containername, in.Podname, in.Namespace, ctx)
	fmt.Printf("fetching log for %v\n", in)
	if err != nil {
		return &ploggi.PodLog{}, err
	}
	return &ploggi.PodLog{Log: logs}, nil
}

func (s *PodLogServer) StreamLog(in *ploggi.Pod, stream ploggi.PodLogs_StreamLogServer) error {
	logs, err := s.LogStreamer.StreamPodLogs(in.Containername, in.Podname, in.Namespace, stream.Context())
	if err != nil {
		return err
	}
	defer logs.Close()
	for {
        buf := make([]byte, 2000)
        numBytes, err := logs.Read(buf)
        if numBytes == 0 {
            continue
        }
        if err == io.EOF {
            break
        }
        if err != nil {
            return err
        }
        message := string(buf[:numBytes])
        stream.Send(&ploggi.PodLog{Log: message})

    }
    return nil
}