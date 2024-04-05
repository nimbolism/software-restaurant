package grpc

import (
	"log"
	"net"

	"github.com/nimbolism/software-restaurant/back-end/user-service/proto"
	"google.golang.org/grpc"
)

type Server struct {
	proto.UnimplementedUserServiceServer
}

func StartServer() {
	println("Starting gRPC server...")
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterUserServiceServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
