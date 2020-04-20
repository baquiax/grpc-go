package main

import (
	"baquiax.me/grpc-go/primes"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

// Server represents a server entity
type Server struct{}

// Calculate streams to the client all the prime factors of the given number
func (s *Server) Calculate(request *primes.Request, server primes.PrimeDecomposition_CalculateServer) error {
	input := request.GetInput()

	var factor int32 = 2

	for input > 1 {
		if input%factor == 0 {
			result := &primes.Result{
				Factor: factor,
			}

			server.Send(result)

			input = input / factor
		} else {
			factor++
		}
	}

	return nil
}

func main() {
	port := 50050
	host := "0.0.0.0"

	log.Printf("Trying to liste on %s:%d\n", host, port)

	listener, error := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))

	if error != nil {
		log.Fatalf("Error trying to listen for connections: %v", error)
	}

	grpcServer := grpc.NewServer()

	primes.RegisterPrimeDecompositionServer(grpcServer, &Server{})

	if error := grpcServer.Serve(listener); error != nil {
		log.Fatalf("Something bad happened serving with gRPC server. %v", error)
	}
}
