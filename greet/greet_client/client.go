package main

import (
	"context"
	"fmt"
	"github.com/fahmyabida/golang-example_grpc/greet/greetpb"
	"google.golang.org/grpc"
	"log"
)

func main(){
	fmt.Println("Hello I'm a client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err !=  nil {
		log.Fatalf("Could NOT connect : %v", err)
	}
	defer cc.Close()
	c := greetpb.NewGreetServiceClient(cc)
	//fmt.Printf("Created client : %f", c)

	req := &greetpb.GreetRequest{
		Greeting:&greetpb.Greeting{
			FirstName: "Fahmy",
			LastName: "Abida",
		},
	}
	//c.Greet(context.Background(), in *greetpb.GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error)
}