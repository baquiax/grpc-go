package main

import (
	"baquiax.me/grpc-go/calculator"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

// Server represents a calculator entity
type Server struct {
}

// Calculate Do a mathematic operation of 2 operands
func (server *Server) Calculate(ctx context.Context, in *calculator.Equation) (*calculator.Result, error) {
	if in.Operation == calculator.Operation_SUM {
		result := calculator.Result{
			Result: in.X + in.Y,
		}

		return &result, nil
	}

	return nil, fmt.Errorf("Unsupported operation %s", in.Operation)
}

func main() {
	port := 5000
	host := "0.0.0.0"

	log.Printf("Trying to listen on %s:%d\n", host, port)

	listener, error := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))

	if error != nil {
		log.Fatalf("Error trying to listen on port %d. %v", port, error)
	}

	grpcServer := grpc.NewServer()

	calculator.RegisterCalculatorServer(grpcServer, &Server{})

	if error := grpcServer.Serve(listener); error != nil {
		log.Fatalf("Error trying to listen on port %d. %v", port, error)
	}
}
