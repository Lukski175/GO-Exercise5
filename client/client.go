package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Lukski175/GO-Exercise5/time"

	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}
	defer conn.Close()

	c := time.NewTimeServiceClient(conn)

	SendTimeRequest(c)
}

func SendTimeRequest(c time.TimeServiceClient) {
	message := time.TimeRequest{}

	response, err := c.GetTime(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling GetTime: %s", err)
	}

	fmt.Printf("Current time right now: %s\n", response.Reply)
}
