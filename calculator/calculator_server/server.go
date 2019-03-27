package main

import (
	"context"
	"fmt"
	"github.com/fahmyabida/golang-example_grpc/calculator/calculatorpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error){
	fmt.Printf("Recieve Sum RPC: %v \n", req)
	firstNumber := req.FirstNumber
	secondNumber := req.SecondNumber
	sum := firstNumber + secondNumber
	res := &calculatorpb.SumResponse{
		SumResult: sum,
	}
	return res,nil
}

func (*server) PrimeNumberDecomposition(req *calculatorpb.PrimeNumberDecompositionRequest, stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error{
	log.Printf("Prime Number Decomposition from number: %v \n", req.Number)
	number := req.GetNumber()
	decompositNumber := int32(2)
	for number > 1 {
		if number%decompositNumber == 0 {
			resultStream := &calculatorpb.PrimeNumberDecompositionResponse{
				ResultDecomposition: decompositNumber,
			}
			stream.Send(resultStream)
			number /= decompositNumber
		} else {
			decompositNumber++
		}
	}
	return nil
}

func main() {
	fmt.Println("Calculator Server")

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Failed to Listen : %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis) ; err != nil {
		log.Fatalf("Failed to Serve : %v", err)
	}
}