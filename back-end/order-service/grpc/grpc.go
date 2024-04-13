package grpc

import (
	"fmt"
	"log"
	"net"

	card_proto "github.com/nimbolism/software-restaurant/back-end/card-service/proto"
	food_proto "github.com/nimbolism/software-restaurant/back-end/food-service/proto"
	"github.com/nimbolism/software-restaurant/back-end/order-service/proto"
	user_proto "github.com/nimbolism/software-restaurant/back-end/user-service/proto"
	voucher_proto "github.com/nimbolism/software-restaurant/back-end/voucher-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	CardServiceClient card_proto.CardServiceClient
	cardClientConn    *grpc.ClientConn

	UserServiceClient user_proto.UserServiceClient
	userClientConn    *grpc.ClientConn

	FoodServiceClient food_proto.FoodServiceClient
	foodClientConn    *grpc.ClientConn

	VoucherServiceClient voucher_proto.VoucherServiceClient
	voucherClientConn    *grpc.ClientConn
)

type Server struct {
	proto.UnimplementedOrderServiceServer
}

func StartServer() {
	println("Starting gRPC server...")
	lis, err := net.Listen("tcp", ":50040")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterOrderServiceServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func InitializeUserGRPCClient() error {
	// Set up a connection to the gRPC server if not already initialized
	if UserServiceClient == nil {
		// Create a connection to the gRPC server
		conn, err := grpc.NewClient("user-service:50010", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return fmt.Errorf("failed to connect to gRPC server: %v", err)
		}

		// Create a client for the UserService
		UserServiceClient = user_proto.NewUserServiceClient(conn)
		userClientConn = conn
	}

	return nil
}

func CloseUserGRPCClient() {
	if userClientConn != nil {
		userClientConn.Close()
	}
}

func InitializeCardGRPCClient() error {
	// Set up a connection to the gRPC server if not already initialized
	if CardServiceClient == nil {
		// Create a connection to the gRPC server
		conn, err := grpc.NewClient("card-service:50020", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return fmt.Errorf("failed to connect to gRPC server: %v", err)
		}

		// Create a client for the UserService
		CardServiceClient = card_proto.NewCardServiceClient(conn)
		cardClientConn = conn
	}

	return nil
}

func CloseCardGRPCClient() {
	if cardClientConn != nil {
		cardClientConn.Close()
	}
}

func InitializeFoodGRPCClient() error {
	// Set up a connection to the gRPC server if not already initialized
	if FoodServiceClient == nil {
		// Create a connection to the gRPC server
		conn, err := grpc.NewClient("food-service:50030", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return fmt.Errorf("failed to connect to gRPC server: %v", err)
		}

		// Create a client for the UserService
		FoodServiceClient = food_proto.NewFoodServiceClient(conn)
		foodClientConn = conn
	}

	return nil
}

func CloseFoodGRPCClient() {
	if foodClientConn != nil {
		foodClientConn.Close()
	}
}

func InitializeVoucherGRPCClient() error {
	// Set up a connection to the gRPC server if not already initialized
	if VoucherServiceClient == nil {
		// Create a connection to the gRPC server
		conn, err := grpc.NewClient("voucher-service:50050", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return fmt.Errorf("failed to connect to gRPC server: %v", err)
		}

		// Create a client for the UserService
		VoucherServiceClient = voucher_proto.NewVoucherServiceClient(conn)
		voucherClientConn = conn
	}

	return nil
}

func CloseVoucherGRPCClient() {
	if voucherClientConn != nil {
		voucherClientConn.Close()
	}
}
