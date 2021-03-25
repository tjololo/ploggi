package server

import (
	"context"

	"github.com/tjololo/ploggi/pkg/api/ploggi"
	"google.golang.org/grpc"
)

type PodLogServer struct {
	ploggi.UnimplementedPodLogsServer
	
}

func (s *PodLogServer) GetLog(ctx context.Context, in *ploggi.Pod, opts ...grpc.CallOption) (*ploggi.PodLog, error) {
	return nil, nil
}