package main

import (
	"context"
	"flag"
	api "grpc-orders_microservice/pkg/api/proto"
	"log"
	"strconv"

	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	if flag.NArg() < 2 {
		log.Fatal("not enough arguments")
	}

	// Correctly handle the error returned by strconv.Atoi
	x, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatalf("failed to convert x to integer: %v", err)
	}

	y, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		log.Fatalf("failed to convert y to integer: %v", err)
	}

	// Establish a connection to the gRPC server
	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial server: %v", err)
	}
	defer conn.Close()

	// Create a new Adder client
	c := api.NewAdderClient(conn)

	// Call the Add method
	res, err := c.Add(context.Background(), &api.AddRequest{X: int32(x), Y: int32(y)})
	if err != nil {
		log.Fatalf("failed to call Add: %v", err)
	}

	// Print the result
	log.Printf("Result: %d", res.GetResult())
}