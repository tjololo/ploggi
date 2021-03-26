package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	pb "github.com/tjololo/ploggi/pkg/api/ploggi"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:8080"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	podName := "ploggi"
	containerName := "ploggi"
	if len(os.Args) > 1 {
		podName = os.Args[1]
	}
	if len(os.Args) > 2 {
		containerName = os.Args[2]
	}
	defer conn.Close()
	c := pb.NewPodLogsClient(conn)

	r, err := c.StreamLog(context.Background(), &pb.Pod{Podname: podName, Namespace: "default", Containername: containerName})
	if err != nil {
		log.Fatalf("could not get log: %v", err)
	}
	log.Printf("Starting to stream messages")
	for {
		in, err := r.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to receive logs: %v", err)
		}
		fmt.Printf("%s", in.Log)
	}
	r.CloseSend()
}