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
	fmt.Print("Received GetTime request\n")
	return &time.TimeReply{Reply: t.Now().String()}, nil
}

func main() {
	fmt.Println("Input port to listen on...")
	var port string
	fmt.Scan(&port)

	// Create listener tcp on port 9080
	fmt.Printf("Listening on port %s", port)
	list, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}
	grpcServer := grpc.NewServer()
	time.RegisterTimeServiceServer(grpcServer, &Server{})

	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("failed to server %v", err)
	}
}
