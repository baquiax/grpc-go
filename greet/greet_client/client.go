package main

import (
	"context"
	"fmt"
	"log"

	"baquiax.me/grpc-go/greet/greetpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Printf("Hello, I'm a client\n\n")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	defer conn.Close()

	if err != nil {
		log.Fatalf("could not connect %v", err)
	}

	client := greetpb.NewGreetServiceClient(conn)

	doUnary(client)
}

func doUnary(client greetpb.GreetServiceClient) {
	fmt.Println("Doing unary request")
	greeting := greetpb.Greeting{
		FirstName: "Alex",
		LastName:  "Baquiax",
	}

	request := greetpb.GreetRequest{
		Greeting: &greeting,
	}

	fmt.Printf("Greeting to server with: %v\n", request)

	response, _ := client.Greet(context.Background(), &request)

	fmt.Printf("Response: %s\n", response.Result)
}
