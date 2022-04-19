package main

import (
	"context"
	"fmt"
	pb "github.com/hedeqiang/grpc-demo/api/simple"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Simple struct {
	Name string
	pb.UnimplementedSimpleServiceServer
}

func (s *Simple) Get(ctx context.Context, req *pb.SimpleRequest) (*pb.SimpleResponse, error) {
	name := req.GetName()
	return &pb.SimpleResponse{
		Message: "Hello " + name,
	}, nil
}

func main() {
	fmt.Println("start server")

	listen, err := net.Listen("tcp", ":8989")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSimpleServiceServer(s, &Simple{})
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
