package main

import (
	"context"
	"fmt"
	"github.com/fahmyabida/golang-example_grpc/calculator/calculatorpb"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main(){

	fmt.Println("Calculator Client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err !=  nil {
		log.Fatalf("Could NOT connect : %v", err)
	}
	defer cc.Close()
	c := calculatorpb.NewCalculatorServiceClient(cc)

	doUnary(c)

	doServerStreaming(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient)  {
	fmt.Println("Starting to do SUM Unary RPC....")
	req := &calculatorpb.SumRequest{
		FirstNumber: 5,
		SecondNumber: 40,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Sum RPC : %v", err)
	}
	log.Printf("Response from Sum: %v", res.SumResult)
}

func doServerStreaming(c calculatorpb.CalculatorServiceClient){
	req := &calculatorpb.PrimeNumberDecompositionRequest{
		Number: 300,
	}
	log.Printf("Prime Number Decomposition : %v", req.GetNumber())
	resStream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling PrimeNumberDecomposition: %v", err)
	}
	for {
		msg, err :=  resStream.Recv()
		if err == io.EOF{
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %", err)
		}
		log.Printf(" %v ", msg.GetResultDecomposition())
	}
}