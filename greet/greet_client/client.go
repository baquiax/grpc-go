package main

import (
	"fmt"
	"log"

	"baquiax.me/grpc-go/greet/greetpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello, I'm a client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	defer conn.Close()

	if err != nil {
		log.Fatalf("could not connect %v", err)
	}

	c := greetpb.NewGreetServiceClient(conn)

	fmt.Printf("Created client %f", c)
}
