package main

import (
	"context"
	"log"
	"os"
	"time"

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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetLog(ctx, &pb.Pod{Podname: podName, Namespace: "default", Containername: containerName})
	if err != nil {
		log.Fatalf("could not get log: %v", err)
	}
	log.Printf("Log: %s", r.Log)
}