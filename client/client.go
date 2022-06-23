package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"google.golang.org/grpc/metadata"
)

// request to hello server
var addr = "localhost:50001"

func request(msg string) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	for {
		func() {
			// Contact the server and print out its response.
			ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
			defer cancel()

			ctx = metadata.AppendToOutgoingContext(ctx, "dapr-app-id", "server")

			r, err := c.SayHello(ctx, &pb.HelloRequest{Name: msg})
			if err != nil {
				log.Printf("could not greet: %+v", err)
			} else {
				log.Printf("Greeting: %s", r.GetMessage())
			}

			time.Sleep(3 * time.Second)
		}()
	}
}
