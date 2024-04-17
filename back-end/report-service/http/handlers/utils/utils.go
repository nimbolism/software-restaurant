package utils

import (
	"context"
	"fmt"

	"github.com/nimbolism/software-restaurant/back-end/card-service/proto"
	"github.com/nimbolism/software-restaurant/back-end/report-service/grpc"
)

func InitializeGRPCClients() error {
	if err := grpc.InitializeUserGRPCClient(); err != nil {
		return fmt.Errorf("failed to initialize user gRPC client: %w", err)
	}
	if err := grpc.InitializeCardGRPCClient(); err != nil {
		return fmt.Errorf("failed to initialize card gRPC client: %w", err)
	}
	if err := grpc.InitializeOrderGRPCClient(); err != nil {
		return fmt.Errorf("failed to initialize order gRPC client: %w", err)
	}
	return nil
}

func GetAccessLevel(cookie string) (int, error) {
	if err := grpc.InitializeCardGRPCClient(); err != nil {
		return 0, fmt.Errorf("failed to initialize card gRPC client: %w", err)
	}
	cardInfo, err := grpc.CardServiceClient.GetCardInfo(context.Background(), &proto.GetCardInfoRequest{JwtToken: cookie})
	if err != nil {
		return 0, fmt.Errorf("failed to retrieve card info: %w", err)
	}
	return int(cardInfo.AccessLevel), nil
}
