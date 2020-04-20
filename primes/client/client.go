package main

import (
	"baquiax.me/grpc-go/primes"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	addr := "0.0.0.0:50050"
	conn, error := grpc.Dial(addr, grpc.WithInsecure())
	defer conn.Close()

	if error != nil {
		log.Fatalf("Error connecting with %s", addr)
	}

	primesClient := primes.NewPrimeDecompositionClient(conn)

	request := &primes.Request{
		Input: 120,
	}

	stream, error := primesClient.Calculate(context.Background(), request)

	for {
		result, error := stream.Recv()

		if error == io.EOF {
			break
		}

		if error != nil {
			log.Fatalf("Error while reading the result. %v", error)
		}

		log.Printf("Factor: %d\n", result.GetFactor())
	}
}
