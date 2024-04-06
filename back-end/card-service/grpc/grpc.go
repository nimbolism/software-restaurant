package grpc

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/nimbolism/software-restaurant/back-end/card-service/http/handlers/utils"
	"github.com/nimbolism/software-restaurant/back-end/card-service/proto"
	"github.com/nimbolism/software-restaurant/back-end/database"
	user_proto "github.com/nimbolism/software-restaurant/back-end/user-service/proto"
)

type Server struct {
	proto.UnimplementedCardServiceServer
}

// Function to get card information
func (s *Server) GetCardInfo(ctx context.Context, req *proto.GetCardInfoRequest) (*proto.CardInfoResponse, error) {
	// Call user-service to authenticate the user
	userServiceClient, err := initializeGRPCClient()
	if err != nil {
		return &proto.CardInfoResponse{
			Error: err.Error(),
		}, nil
	}

	// Call the AuthenticateUser function of the user-service
	authenticateUserResponse, err := userServiceClient.AuthenticateUser(ctx, &user_proto.AuthenticateUserRequest{JwtToken: req.JwtToken})
	if err != nil {
		return &proto.CardInfoResponse{
			Error: err.Error(),
		}, nil
	}

	// If authentication fails, populate the error field and return
	if !authenticateUserResponse.Success {
		return &proto.CardInfoResponse{
			Error: authenticateUserResponse.ErrorMessage,
		}, nil
	}

	db := database.GetPQDB()
	card, err := utils.FindCardByUserID(db, uint(authenticateUserResponse.UserId))
	if err != nil {
		return &proto.CardInfoResponse{
			Error: err.Error(),
		}, nil
	}

	// Dummy implementation: Return hardcoded card information
	return &proto.CardInfoResponse{
		BlackListed: card.BlackListed,
		AccessLevel: int32(card.AccessLevel),
	}, nil
}

// Function to update reserves count
func (s *Server) UpdateReserves(ctx context.Context, req *proto.UpdateReservesRequest) (*empty.Empty, error) {
	// Call user-service to authenticate the user
	userServiceClient, err := initializeGRPCClient()
	if err != nil {
		log.Fatalf("Failed to initialize gRPC client: %v", err)
	}

	// Call the AuthenticateUser function of the user-service
	authenticateUserResponse, err := userServiceClient.AuthenticateUser(ctx, &user_proto.AuthenticateUserRequest{JwtToken: req.JwtToken})
	if err != nil {
		return nil, err
	}

	// If authentication fails, return appropriate error
	if !authenticateUserResponse.Success {
		return nil, errors.New(authenticateUserResponse.ErrorMessage)
	}

	db := database.GetPQDB()
	// Find the card in the database by UserID
	card, err := utils.FindCardByUserID(db, uint(authenticateUserResponse.UserId))
	if err != nil {
		return nil, err
	}

	// Update reserves count
	card.Reserves += int(req.ReservesChange)
	if err := db.Save(&card).Error; err != nil {
		return nil, err
	}

	// Return empty response
	return &empty.Empty{}, nil
}

func StartServer() {
	println("Starting gRPC server...")
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterCardServiceServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func initializeGRPCClient() (user_proto.UserServiceClient, error) {
	// Set up a connection to the gRPC server
	conn, err := grpc.Dial("localhost:50050", grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Create a client for the CardService
	client := user_proto.NewUserServiceClient(conn)
	return client, nil
}
