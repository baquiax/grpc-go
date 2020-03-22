package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"baquiax.me/grpc-go/greet/greetpb"
	"google.golang.org/grpc"
)

// Server struct of a very simple gRPC server
type Server struct{}

// Greet Simple greeting function
func (s *Server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	firstName := req.GetGreeting().GetFirstName()
	fmt.Printf("A greet was received %v", req.GetGreeting())

	result := greetpb.GreetResponse{
		Result: fmt.Sprintf("Hello, %s", firstName),
	}

	return &result, nil
}

func main() {
	host := "0.0.0.0"
	port := "50051"

	fmt.Printf("Starting a new server at: %s:%s\n\n", host, port)

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))

	if err != nil {
		log.Fatalf("Error listening on %s:%s", host, port)
	}

	server := grpc.NewServer()

	greetpb.RegisterGreetServiceServer(server, &Server{})

	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
