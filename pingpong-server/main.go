package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "pingpong-server/pb"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (*server) PingPong(ctx context.Context, req *pb.PingRequest) (*pb.PongResponse, error) {
	fmt.Println("Call PingPong Server...")
	res := &pb.PongResponse{
		Pong: "Pong pong Pong",
	}
	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	fmt.Println("Pingpong server start...")

	opts := []grpc.ServerOption{}

	s := grpc.NewServer(opts...)
	pb.RegisterPingPongServiceServer(s, &server{})

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
