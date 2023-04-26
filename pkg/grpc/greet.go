package grpc

import (
	"context"
	"fmt"
	"github.com/yinzhenzhen/go-grpc-demo/pkg/greet"
	"log"
)

type GreetServer struct{}

func (s *GreetServer) Greet(ctx context.Context, req *greet.GreetRequest) (*greet.GreetResponse, error) {
	return &greet.GreetResponse{Result: "Hello " + req.Greeting.FirstName}, nil
}

//func (s *GreetServer) GreetManyTimes(req *greet.GreetManyTimesRequest) (*greet.GreetManyTimesResponse, error) {
//	return &greet.GreetManyTimesResponse{Result: req.Greeting.FirstName}, nil
//}
//
//func (s *GreetServer) LongGreet(req *greet.LongGreetRequest) (*greet.LongGreetResponse, error) {
//	return &greet.LongGreetResponse{Result: req.Greeting.FirstName}, nil
//}
//
//func (s *GreetServer) GreetEveryone(req *greet.GreetEveryoneRequest) (*greet.GreetEveryoneResponse, error) {
//	return &greet.GreetEveryoneResponse{Result: req.Greeting.FirstName}, nil
//}
//
//func (s *GreetServer) GreetWithDeadline(ctx context.Context, req *greet.GreetWithDeadlineRequest) (*greet.GreetWithDeadlineResponse, error) {
//	return &greet.GreetWithDeadlineResponse{Result: req.Greeting.FirstName}, nil
//}

func greetApi(c greet.GreetServiceClient) {
	r, err := c.Greet(context.Background(), &greet.GreetRequest{Greeting: &greet.Greeting{FirstName: "zhen", LastName: "yin"}})
	if err != nil {
		log.Fatalf("failed to greet: %v", err)
	} else {
		fmt.Printf("%s\n", r.Result)
	}
}
