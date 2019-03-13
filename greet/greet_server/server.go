package main

import (
	"context"
	"fmt"
	"github.com/fahmyabida/golang-example_grpc/greet/greetpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error){
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello "+ firstName
	res := &greetpb.GreetResponse{
		Result : result,
	}
	return res, nil
}

func main() {
	fmt.Println("Hello World")

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Failed to Listen : %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis) ; err != nil {
		log.Fatalf("Failed to Serve : %v", err)
	}
}