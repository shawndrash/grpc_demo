package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/shawndrash/grpc_demo/proto"
	"google.golang.org/grpc"
)

type greeterServer struct {
	pb.UnimplementedGreeterServer
}

func (g *greeterServer) Greet(ctx context.Context, req *pb.GreetingRequest) (*pb.GreetingResponse, error) {
	return &pb.GreetingResponse{
		Greeting: "Hello " + req.Name,
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &greeterServer{})
	fmt.Println("Server started on port 50051")
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
