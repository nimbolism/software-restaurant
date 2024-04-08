package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	card_proto "github.com/nimbolism/software-restaurant/back-end/card-service/proto"
	"github.com/nimbolism/software-restaurant/back-end/database"
	"github.com/nimbolism/software-restaurant/back-end/database/models"
	"github.com/nimbolism/software-restaurant/back-end/food-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	CardServiceClient card_proto.CardServiceClient
	cardClientConn    *grpc.ClientConn
)

type Server struct {
	proto.UnimplementedFoodServiceServer
}

func (s *Server) GetFoodDetailsById(ctx context.Context, req *proto.FoodIdRequest) (*proto.Food, error) {
	db := database.GetPQDB()
	var food models.Food
	if err := db.First(&food, req.FoodId).Error; err != nil {
		return &proto.Food{
			ErrorMessage: fmt.Sprintf("Food with ID %d not found: %v", req.FoodId, err),
		}, nil
	}

	var meal models.Meal
	if err := db.First(&meal, food.MealID).Error; err != nil {
		return &proto.Food{
			ErrorMessage: fmt.Sprintf("meal with ID %d not found: %v", food.MealID, err),
		}, nil
	}
	var category models.Category
	if err := db.First(&category, food.CategoryID).Error; err != nil {
		return &proto.Food{
			ErrorMessage: fmt.Sprintf("category with ID %d not found: %v", food.CategoryID, err),
		}, nil
	}

	// Convert database food model to protobuf message
	foodPB := &proto.Food{
		Id:          int64(food.ID),
		Name:        food.Name,
		Description: food.Description,
		Category:    category.Name,
		Meal:        meal.Name,
	}

	createdAtProto := timestamppb.New(food.CreatedAt)
	updatedAtProto := timestamppb.New(food.UpdatedAt)
	foodPB.CreatedAt = createdAtProto
	foodPB.UpdatedAt = updatedAtProto

	return foodPB, nil
}

func (s *Server) GetSideDishDetailsById(ctx context.Context, req *proto.SideDishIdRequest) (*proto.SideDish, error) {
	db := database.GetPQDB()
	var sideDish models.SideDish
	if err := db.First(&sideDish, req.SideDishId).Error; err != nil {
		return &proto.SideDish{
			ErrorMessage: fmt.Sprintf("Food with ID %d not found: %v", req.SideDishId, err),
		}, nil
	}

	// Convert database side dish model to protobuf message
	sideDishPB := &proto.SideDish{
		Id:          int64(sideDish.ID),
		Name:        sideDish.Name,
		Description: sideDish.Description,
	}
	createdAtProto := timestamppb.New(sideDish.CreatedAt)
	updatedAtProto := timestamppb.New(sideDish.UpdatedAt)
	sideDishPB.CreatedAt = createdAtProto
	sideDishPB.UpdatedAt = updatedAtProto

	return sideDishPB, nil
}

func StartServer() {
	println("Starting gRPC server...")
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterFoodServiceServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func InitializeGRPCClient() error {
	// Set up a connection to the gRPC server if not already initialized
	if CardServiceClient == nil {
		// Create a connection to the gRPC server
		conn, err := grpc.NewClient("card-service:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return fmt.Errorf("failed to connect to gRPC server: %v", err)
		}

		// Create a client for the UserService
		CardServiceClient = card_proto.NewCardServiceClient(conn)
		cardClientConn = conn
	}

	return nil
}

func CloseGRPCClient() {
	if cardClientConn != nil {
		cardClientConn.Close()
	}
}
