package main

import (
	"context"
	"fmt"
	"log"
	"net"
	t "time"

	"github.com/Lukski175/GO-Exercise5/time"

	"google.golang.org/grpc"
)

type Server struct {
	time.UnimplementedTimeServiceServer
}

func (s *Server) GetTime(ctx context.Context, in *time.TimeRequest) (*time.TimeReply, error) {
	fmt.Printf("Received GetTime request\n")
	return &time.GetTimeReply{Reply: t.Now().String()}, nil
}

func main() {
	// Create listener tcp on port 9080
	list, err := net.Listen("tcp", ":9080")
	if err != nil {
		log.Fatalf("Failed to listen on port 9080: %v", err)
	}
	grpcServer := grpc.NewServer()
	time.RegisterTimeServiceServer(grpcServer, &Server{})

	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("failed to server %v", err)
	}
}
