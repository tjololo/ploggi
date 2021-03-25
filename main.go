package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/tjololo/ploggi/pkg/api/ploggi"
	"github.com/tjololo/ploggi/pkg/server"
	"google.golang.org/grpc"
	//"k8s.io/client-go/kubernetes"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 50051, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPodLogsServer(s, &server.PodLogServer{})
	fmt.Printf("starting server on port %d\n", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
