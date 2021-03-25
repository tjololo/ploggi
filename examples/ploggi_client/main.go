package main

import (
	"context"
	"log"
	"time"

	pb "github.com/tjololo/ploggi/pkg/api/ploggi"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewPodLogsClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetLog(ctx, &pb.Pod{Podname: "test", Namespace: "default", Containername: "deployment"})
	if err != nil {
		log.Fatalf("could not get log: %v", err)
	}
	log.Printf("Log: %s", r.Log)
}