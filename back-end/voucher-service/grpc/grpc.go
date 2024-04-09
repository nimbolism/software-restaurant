package grpc

import (
	"log"
	"net"

	"github.com/nimbolism/software-restaurant/back-end/voucher-service/proto"
	"google.golang.org/grpc"
)

type Server struct {
	proto.UnimplementedVoucherServiceServer
}

func StartServer() {
	println("Starting gRPC server...")
	lis, err := net.Listen("tcp", ":50054")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterVoucherServiceServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
