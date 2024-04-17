package grpc

import (
	"context"
	"log"
	"net"

	"github.com/nimbolism/software-restaurant/back-end/database/models"
	"github.com/nimbolism/software-restaurant/back-end/gutils/postgresapp"
	"github.com/nimbolism/software-restaurant/back-end/user-service/http/handlers/auth"
	"github.com/nimbolism/software-restaurant/back-end/user-service/http/handlers/utils"
	"github.com/nimbolism/software-restaurant/back-end/user-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	proto.UnimplementedUserServiceServer
}

func (s *Server) AuthenticateUser(ctx context.Context, req *proto.AuthenticateUserRequest) (*proto.AuthenticateUserResponse, error) {
	cookie := req.JwtToken
	username, err := auth.GetUsernameFromJWT(cookie)
	if err != nil {
		return nil, status.Error(
			codes.PermissionDenied, "Invalid credentials",
		)
	}

	db := postgresapp.DB
	existingUser, err := utils.GetExistingUser(db, username)
	if err != nil {
		return nil, status.Error(
			codes.NotFound, "Such user does not exist",
		)
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
		return nil, status.Error(
			codes.PermissionDenied, "Invalid credentials",
		)
	}

	db := postgresapp.DB
	existingUser, err := utils.GetExistingUser(db, username)
	if err != nil {
		return nil, status.Error(
			codes.NotFound, "Such user does not exist",
		)
	}

	return &proto.GetUserInfoResponse{
		UserData: &proto.UserData{
			UserId:       uint64(existingUser.ID),
			Username:     existingUser.Username,
			Email:        existingUser.Email,
			PhoneNumber:  existingUser.PhoneNumber,
			NationalCode: existingUser.NationalCode,
		},
	}, nil
}

func (s *Server) GetAllUsers(ctx context.Context, req *proto.GetAllUsersRequest) (*proto.GetAllUsersResponse, error) {
	var users []models.User
	db := postgresapp.DB
	if err := db.Find(&users).Error; err != nil {
		return nil, status.Error(
			codes.Internal, "Cannot get users from database",
		)
	}

	var userData []*proto.UserData
	for _, user := range users {
		userData = append(userData, &proto.UserData{
			Username:     user.Username,
			Email:        user.Email,
			PhoneNumber:  user.PhoneNumber,
			NationalCode: user.NationalCode,
		})
	}

	return &proto.GetAllUsersResponse{Users: userData}, nil
}

func (s *Server) GetOneUser(ctx context.Context, req *proto.GetOneUserRequest) (*proto.GetOneUserResponse, error) {
	db := postgresapp.DB
	reqUser, err := utils.GetExistingUser(db, req.Username)
	if err != nil {
		return nil, status.Error(
			codes.NotFound, "Such user does not exist",
		)
	}
	return &proto.GetOneUserResponse{
		UserId: uint64(reqUser.ID),
	}, nil
}

func StartServer() {
	println("Starting gRPC server...")
	lis, err := net.Listen("tcp", ":50010")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterUserServiceServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
