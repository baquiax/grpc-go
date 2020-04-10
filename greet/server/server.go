package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"baquiax.me/grpc-go/greet/greetpb"
	"google.golang.org/grpc"
)

// Server struct of a very simple gRPC server
type Server struct{}

// Greet Simple greeting function
func (s *Server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	firstName := req.GetGreeting().GetFirstName()
	log.Printf("A greet request was received with %v", req.GetGreeting())

	result := greetpb.GreetResponse{
		Result: fmt.Sprintf("Hello, %s", firstName),
	}

	return &result, nil
}

// GreetManyTimes stream messages from the server
func (s *Server) GreetManyTimes(req *greetpb.GreetRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	log.Printf("A greet many times request was received with %v", req.GetGreeting())
	firstName := req.GetGreeting().GetFirstName()

	for i := 1; i <= 10; i++ {
		result := fmt.Sprintf("Greet #%d for %s", i, firstName)
		res := &greetpb.GreetResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

func main() {
	host := "0.0.0.0"
	port := "50051"

	log.Printf("Starting a new server at: %s:%s\n\n", host, port)

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
