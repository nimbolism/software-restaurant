package grpc

import (
	"context"
	"log"
	"net"

	"github.com/nimbolism/software-restaurant/back-end/database"
	"github.com/nimbolism/software-restaurant/back-end/database/models"
	"github.com/nimbolism/software-restaurant/back-end/user-service/http/handlers/auth"
	"github.com/nimbolism/software-restaurant/back-end/user-service/http/handlers/utils"
	"github.com/nimbolism/software-restaurant/back-end/user-service/proto"
	"google.golang.org/grpc"
)

type Server struct {
	proto.UnimplementedUserServiceServer
}

func (s *Server) AuthenticateUser(ctx context.Context, req *proto.AuthenticateUserRequest) (*proto.AuthenticateUserResponse, error) {
	cookie := req.JwtToken
	username, err := auth.GetUsernameFromJWT(cookie)
	if err != nil {
		return &proto.AuthenticateUserResponse{
			Success:      false,
			ErrorMessage: "Invalid credentials",
		}, nil
	}

	db := database.GetPQDB()
	existingUser, err := utils.GetExistingUser(db, username)
	if err != nil {
		return &proto.AuthenticateUserResponse{
			Success:      false,
			ErrorMessage: "Invalid credentials ",
		}, nil
	}

	return &proto.AuthenticateUserResponse{
		UserId:  uint64(existingUser.ID),
		Success: true,
	}, nil
}

func (s *Server) GetUserInfo(ctx context.Context, req *proto.GetUserInfoRequest) (*proto.GetUserInfoResponse, error) {
	cookie := req.JwtToken
	username, err := auth.GetUsernameFromJWT(cookie)
	if err != nil {
		return &proto.GetUserInfoResponse{
			ErrorMessage: "Invalid credentials",
		}, nil
	}

	db := database.GetPQDB()
	existingUser, err := utils.GetExistingUser(db, username)
	if err != nil {
		return &proto.GetUserInfoResponse{
			ErrorMessage: "Invalid credentials",
		}, nil
	}

	return &proto.GetUserInfoResponse{
		Username:     existingUser.Username,
		Email:        existingUser.Email,
		PhoneNumber:  existingUser.PhoneNumber,
		NationalCode: existingUser.PhoneNumber,
	}, nil
}

func (s *Server) GetAllUsers(ctx context.Context, req *proto.GetAllUsersRequest) (*proto.GetAllUsersResponse, error) {
	var users []models.User
	db := database.GetPQDB()
	if err := db.Find(&users).Error; err != nil {
		return &proto.GetAllUsersResponse{
			ErrorMessage: "Error in database",
		}, nil
	}

	var userData []*proto.UserData
	for _, user := range userData {
		userData = append(userData, &proto.UserData{
			Username:     user.Username,
			Email:        user.Email,
			PhoneNumber:  user.PhoneNumber,
			NationalCode: user.NationalCode,
		})
	}

	return &proto.GetAllUsersResponse{Users: userData}, nil
}

func StartServer() {
	println("Starting gRPC server...")
	lis, err := net.Listen("tcp", ":50050")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterUserServiceServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
