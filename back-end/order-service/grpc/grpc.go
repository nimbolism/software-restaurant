package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	card_proto "github.com/nimbolism/software-restaurant/back-end/card-service/proto"
	"github.com/nimbolism/software-restaurant/back-end/database/models"
	food_proto "github.com/nimbolism/software-restaurant/back-end/food-service/proto"
	"github.com/nimbolism/software-restaurant/back-end/gutils/postgresapp"
	"github.com/nimbolism/software-restaurant/back-end/order-service/proto"
	user_proto "github.com/nimbolism/software-restaurant/back-end/user-service/proto"
	voucher_proto "github.com/nimbolism/software-restaurant/back-end/voucher-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
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

func (s *Server) GetAllOrders(ctx context.Context, req *proto.GetAllOrdersRequest) (*proto.GetAllOrdersResponse, error) {
	db := postgresapp.DB
	var orders []models.Order
	if err := db.Find(&orders).Error; err != nil {
		return nil, status.Error(
			codes.NotFound, fmt.Sprintf("could not find orders: %v", err),
		)
	}
	return &proto.GetAllOrdersResponse{Orders: convertToProtoOrders(orders)}, nil
}

func (s *Server) GetAllOrdersByUsername(ctx context.Context, req *proto.GetAllOrdersByUsernameRequest) (*proto.GetAllOrdersResponse, error) {
	err := InitializeUserGRPCClient()
	if err != nil {
		return nil, status.Error(
			codes.NotFound, fmt.Sprintf("could not initialize connection: %v", err),
		)
	}
	userIDResponse, err := UserServiceClient.GetOneUser(context.Background(), &user_proto.GetOneUserRequest{Username: req.Username})
	if err != nil {
		return nil, status.Error(
			codes.NotFound, fmt.Sprintf("could not get response from user service: %v", err),
		)
	}
	db := postgresapp.DB
	var orders []models.Order
	if err := db.Where("user_id = ?", userIDResponse.UserId).Find(&orders).Error; err != nil {
		return nil, status.Error(
			codes.NotFound, fmt.Sprintf("could not get orders from database: %v", err),
		)
	}
	return &proto.GetAllOrdersResponse{Orders: convertToProtoOrders(orders)}, nil
}

func (s *Server) GetAllOrdersBetweenTimestamps(ctx context.Context, req *proto.GetAllOrdersBetweenTimestampsRequest) (*proto.GetAllOrdersResponse, error) {
	db := postgresapp.DB
	var orders []models.Order
	if err := db.Where("created_at BETWEEN ? AND ?", req.StartTime.AsTime(), req.EndTime.AsTime()).Find(&orders).Error; err != nil {
		return nil, status.Error(
			codes.NotFound, fmt.Sprintf("could not find orders: %v", err),
		)
	}
	return &proto.GetAllOrdersResponse{Orders: convertToProtoOrders(orders)}, nil
}

func convertToProtoOrders(orders []models.Order) []*proto.Order {
	protoOrders := make([]*proto.Order, len(orders))
	for i, order := range orders {
		protoOrder := &proto.Order{
			Id:     uint64(order.ID),
			UserId: uint64(order.UserID),
			Paid:   order.Paid,
			CreatedAt: &timestamppb.Timestamp{
				Seconds: order.CreatedAt.Unix(),
				Nanos:   int32(order.CreatedAt.Nanosecond()),
			},
			UpdatedAt: &timestamppb.Timestamp{
				Seconds: order.UpdatedAt.Unix(),
				Nanos:   int32(order.UpdatedAt.Nanosecond()),
			},
		}

		protoFoods := make([]*proto.OrderFood, len(order.Foods))
		for j, food := range order.Foods {
			protoFoods[j] = &proto.OrderFood{
				Id:          int64(food.ID),
				Name:        food.Name,
				Description: food.Description,
			}
		}
		protoOrder.Foods = protoFoods

		protoSideDishes := make([]*proto.OrderSideDish, len(order.SideDishes))
		for k, sideDish := range order.SideDishes {
			protoSideDishes[k] = &proto.OrderSideDish{
				Id:          int64(sideDish.ID),
				Name:        sideDish.Name,
				Description: sideDish.Description,
			}
		}
		protoOrder.SideDishes = protoSideDishes

		protoOrders[i] = protoOrder
	}
	return protoOrders
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
