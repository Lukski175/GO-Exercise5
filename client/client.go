package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Lukski175/GO-Exercise5/time"

	"google.golang.org/grpc"
)

func main() {

	s1, c1 := SetupServer()
	s2, c2 := SetupServer()
	defer c1.Close()
	defer c2.Close()

	//fmt.Printf("After setup: \nOne: %d \n Two:%e \n", s1, s2)

	go SendTimeRequest(s1)
	go SendTimeRequest(s2)

	for {

	}
}

func SetupServer() (time.TimeServiceClient, grpc.ClientConn) {
	fmt.Println("Input port to connect to...")
	var port string
	fmt.Scan(&port)
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("0.0.0.0:"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}

	c := time.NewTimeServiceClient(conn)
	//fmt.Printf("Setup server service: %d \n", s)

	return c, *conn
}

func SendTimeRequest(c time.TimeServiceClient) {
	message := time.TimeRequest{}

	response, err := c.GetTime(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling GetTime: %s", err)
	}

	fmt.Printf("Current time right now: %s\n", response.Reply)
}
