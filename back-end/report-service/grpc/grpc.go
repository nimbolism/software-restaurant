package grpc

import (
	"fmt"

	card_proto "github.com/nimbolism/software-restaurant/back-end/card-service/proto"
	order_proto "github.com/nimbolism/software-restaurant/back-end/order-service/proto"
	user_proto "github.com/nimbolism/software-restaurant/back-end/user-service/proto"
	voucher_proto "github.com/nimbolism/software-restaurant/back-end/voucher-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	UserServiceClient user_proto.UserServiceClient
	userClientConn    *grpc.ClientConn

	CardServiceClient card_proto.CardServiceClient
	cardClientConn    *grpc.ClientConn

	OrderServiceClient order_proto.OrderServiceClient
	orderClientConn    *grpc.ClientConn

	VoucherServiceClient voucher_proto.VoucherServiceClient
	voucherClientConn    *grpc.ClientConn
)

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

func InitializeOrderGRPCClient() error {
	// Set up a connection to the gRPC server if not already initialized
	if OrderServiceClient == nil {
		// Create a connection to the gRPC server
		conn, err := grpc.NewClient("order-service:50040", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return fmt.Errorf("failed to connect to gRPC server: %v", err)
		}

		// Create a client for the UserService
		OrderServiceClient = order_proto.NewOrderServiceClient(conn)
		orderClientConn = conn
	}

	return nil
}

func CloseOrderGRPCClient() {
	if orderClientConn != nil {
		orderClientConn.Close()
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
