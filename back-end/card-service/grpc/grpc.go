package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/nimbolism/software-restaurant/back-end/card-service/http/handlers/utils"
	"github.com/nimbolism/software-restaurant/back-end/card-service/proto"
	"github.com/nimbolism/software-restaurant/back-end/gutils"
	"github.com/nimbolism/software-restaurant/back-end/gutils/postgresapp"
	user_proto "github.com/nimbolism/software-restaurant/back-end/user-service/proto"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	UserServiceClient user_proto.UserServiceClient
	userClientConn    *grpc.ClientConn
)

type Server struct {
	proto.UnimplementedCardServiceServer
}

// Function to get card information
func (s *Server) GetCardInfo(ctx context.Context, req *proto.GetCardInfoRequest) (*proto.CardInfoResponse, error) {
	authenticateUserResponse, err := AuthenticateUserService(ctx, req.JwtToken)
	if err != nil {
		return nil, err
	}

	card, err := utils.FindCardByUserID(uint(authenticateUserResponse.UserId))
	if err != nil {
		return nil, err
	}

	return &proto.CardInfoResponse{
		BlackListed: card.BlackListed,
		Verified:    card.Verified,
		AccessLevel: int32(card.AccessLevel),
	}, nil
}

func (s *Server) UpdateReserves(ctx context.Context, req *proto.UpdateReservesRequest) (*empty.Empty, error) {
	authenticateUserResponse, err := AuthenticateUserService(ctx, req.JwtToken)
	if err != nil {
		return nil, err
	}

	card, err := utils.FindCardByUserID(uint(authenticateUserResponse.UserId))
	if err != nil {
		return nil, err
	}

	db := postgresapp.DB
	card.Reserves += int(req.ReservesChange)
	if err := db.Save(&card).Error; err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func AuthenticateUserService(ctx context.Context, jwtToken string) (*user_proto.AuthenticateUserResponse, error) {
	if err := InitializeGRPCClient(); err != nil {
		return nil, fmt.Errorf("failed to initialize gRPC client: %v", err)
	}

	return UserServiceClient.AuthenticateUser(ctx, &user_proto.AuthenticateUserRequest{JwtToken: jwtToken})
}

func AuthenticateUser(c *fiber.Ctx) (*user_proto.AuthenticateUserResponse, error) {
	cookie := gutils.GetCookie(c, "jwt")
	if cookie == "" {
		return nil, c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "JWT cookie not found"})
	}

	if err := InitializeGRPCClient(); err != nil {
		log.Fatalf("Failed to initialize gRPC client: %v", err)
	}

	return UserServiceClient.AuthenticateUser(context.Background(), &user_proto.AuthenticateUserRequest{JwtToken: cookie})
}

func StartServer() {
	println("Starting gRPC server...")
	lis, err := net.Listen("tcp", ":50020")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterCardServiceServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func InitializeGRPCClient() error {
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

func CloseGRPCClient() {
	if userClientConn != nil {
		userClientConn.Close()
	}
}
