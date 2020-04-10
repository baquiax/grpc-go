package main

import (
	"context"
	"io"
	"log"

	"baquiax.me/grpc-go/greet/greetpb"

	"google.golang.org/grpc"
)

func main() {
	log.Printf("Hello, I'm a client\n\n")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	defer conn.Close()

	if err != nil {
		log.Fatalf("could not connect %v", err)
	}

	client := greetpb.NewGreetServiceClient(conn)

	// doUnary(client)

	doServerStream(client)
}

func doUnary(client greetpb.GreetServiceClient) {
	log.Println("Doing unary request")
	greeting := greetpb.Greeting{
		FirstName: "Alex",
		LastName:  "Baquiax",
	}

	request := greetpb.GreetRequest{
		Greeting: &greeting,
	}

	log.Printf("Greeting to server with: %v\n", request)

	response, _ := client.Greet(context.Background(), &request)

	log.Printf("Response: %s\n", response.Result)
}

func doServerStream(client greetpb.GreetServiceClient) {
	log.Println("Doing server stream request")

	request := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Alex",
			LastName:  "Baquiax",
		},
	}

	streamResponse, error := client.GreetManyTimes(context.Background(), request)

	if error != nil {
		log.Fatalf("An error happened during the request %v", error)
	}

	for {
		response, error := streamResponse.Recv()

		if error == io.EOF {
			return
		}

		if error != nil {
			log.Fatalf("Something happened during the lecture of the response: %v", error)
		}

		log.Println(response.Result)
	}
}
