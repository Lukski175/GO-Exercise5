package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/Lukski175/GO-Exercise5/time"

	"google.golang.org/grpc"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	var s1 time.TimeServiceClient
	var s2 time.TimeServiceClient
	go SetupServer(&wg, &s1)
	go SetupServer(&wg, &s2)
	wg.Wait()

	go SendTimeRequest(s1)
	go SendTimeRequest(s2)
}

func SetupServer(gr *sync.WaitGroup, s *time.TimeServiceClient) {
	fmt.Println("Input port to connect to...")
	var port string
	fmt.Scan(&port)
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}
	defer conn.Close()

	c := time.NewTimeServiceClient(conn)
	s = &c

	gr.Done()
}

func SendTimeRequest(c time.TimeServiceClient) {
	message := time.TimeRequest{}

	response, err := c.GetTime(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling GetTime: %s", err)
	}

	fmt.Printf("Current time right now: %s\n", response.Reply)
}
