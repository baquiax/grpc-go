package main

import (
	"baquiax.me/grpc-go/calculator"
	"context"
	"google.golang.org/grpc"
	"log"
)

func main() {
	addr := "0.0.0.0:5000"
	conn, error := grpc.Dial(addr, grpc.WithInsecure())
	defer conn.Close()

	if error != nil {
		log.Fatalf("Not able to open connection to %s", addr)
	}

	calcClient := calculator.NewCalculatorClient(conn)

	equation := &calculator.Equation{
		Operation: calculator.Operation_SUM,
		X:         10,
		Y:         10,
	}

	resp, error := calcClient.Calculate(context.Background(), equation)

	if error != nil {
		log.Fatalf("An error occured when trying to solve an equation %v", error)
	}

	log.Printf("The operation: %s(%d, %d) = %d", equation.Operation.String(), equation.X, equation.Y, resp.Result)
}
