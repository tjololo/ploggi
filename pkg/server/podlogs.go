package server

import (
	"context"

	"github.com/tjololo/ploggi/pkg/api/ploggi"
	plogs "github.com/tjololo/ploggi/pkg/logs"
)

type PodLogServer struct {
	ploggi.UnimplementedPodLogsServer
	LogStreamer *plogs.Streamer
}

func (s *PodLogServer) GetLog(ctx context.Context, in *ploggi.Pod) (*ploggi.PodLog, error) {
	logs, err := s.LogStreamer.PodLogs(in.Containername, in.Podname, in.Namespace, ctx)
	if err != nil {
		return &ploggi.PodLog{}, err
	}
	return &ploggi.PodLog{Log: logs}, nil
}