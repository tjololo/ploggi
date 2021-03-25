package server

import (
	"context"

	"github.com/tjololo/ploggi/pkg/api/ploggi"
)

type PodLogServer struct {
	ploggi.UnimplementedPodLogsServer
	
}

func (s *PodLogServer) GetLog(ctx context.Context, in *ploggi.Pod) (*ploggi.PodLog, error) {
	return &ploggi.PodLog{Log: "this would be the logs"}, nil
}