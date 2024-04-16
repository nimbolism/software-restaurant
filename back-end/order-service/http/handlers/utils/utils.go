package utils

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"
	card_proto "github.com/nimbolism/software-restaurant/back-end/card-service/proto"
	"github.com/nimbolism/software-restaurant/back-end/database/models"
	food_proto "github.com/nimbolism/software-restaurant/back-end/food-service/proto"
	"github.com/nimbolism/software-restaurant/back-end/order-service/grpc"
	user_proto "github.com/nimbolism/software-restaurant/back-end/user-service/proto"
	voucher_proto "github.com/nimbolism/software-restaurant/back-end/voucher-service/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

func SaveOrder(order *models.Order, db *gorm.DB) error {
	if grpc.VoucherServiceClient == nil {
		if err := grpc.InitializeVoucherGRPCClient(); err != nil {
			return fmt.Errorf("failed to initialize voucher gRPC client: %w", err)
		}
	}
	fmt.Println(convertOrderToProto(order))
	storeResponse, err := grpc.VoucherServiceClient.StoreOrderDetails(context.Background(), &voucher_proto.StoreOrderDetailsRequestHelper{Order: convertOrderToProto(order)})
	if err != nil {
		return fmt.Errorf("failed to store order details: %v", err)
	}

	fmt.Println("after function")
	// Handle response
	if !storeResponse.GetSuccess() {
		return errors.New("failed to store order details: response indicates failure")
	}

	return nil
}

func convertOrderToProto(order *models.Order) *voucher_proto.OrderHelper {
	// Convert foods
	protoFoods := make([]*voucher_proto.FoodHelper, len(order.Foods))
	for i, food := range order.Foods {
		protoFoods[i] = &voucher_proto.FoodHelper{
			Name:         food.Name,
			Description:  food.Description,
			CategoryName: food.CategoryName,
			MealName:     food.MealName,
		}
	}

	// Convert side dishes
	protoSideDishes := make([]*voucher_proto.SideDishHelper, len(order.SideDishes))
	for i, sideDish := range order.SideDishes {
		protoSideDishes[i] = &voucher_proto.SideDishHelper{
			Name:        sideDish.Name,
			Description: sideDish.Description,
		}
	}

	// Convert order
	protoOrder := &voucher_proto.OrderHelper{
		Id:          uint64(order.ID), // Assuming ID is uint
		Username:    order.User.Username,
		Email:       order.User.Email,
		PhoneNumber: order.User.PhoneNumber,
		Foods:       protoFoods,
		SideDishes:  protoSideDishes,
		Paid:        order.Paid,
		CreatedAt:   timestamppb.New(order.CreatedAt), // Assuming CreatedAt is of type time.Time in models.Order
		UpdatedAt:   timestamppb.New(order.UpdatedAt), // Assuming UpdatedAt is of type time.Time in models.Order
	}

	return protoOrder
}

func InitializeGRPCClients() error {
	if err := grpc.InitializeUserGRPCClient(); err != nil {
		return fmt.Errorf("failed to initialize user gRPC client: %w", err)
	}
	if err := grpc.InitializeCardGRPCClient(); err != nil {
		return fmt.Errorf("failed to initialize card gRPC client: %w", err)
	}
	if err := grpc.InitializeFoodGRPCClient(); err != nil {
		return fmt.Errorf("failed to initialize food gRPC client: %w", err)
	}
	if err := grpc.InitializeVoucherGRPCClient(); err != nil {
		return fmt.Errorf("failed to initialize voucher gRPC client: %w", err)
	}
	return nil
}

func GetCardDetails(cookie string) (*card_proto.CardInfoResponse, error) {
	cardDetails, err := grpc.CardServiceClient.GetCardInfo(context.Background(), &card_proto.GetCardInfoRequest{JwtToken: cookie})
	if err != nil {
		return nil, err
	}
	return cardDetails, nil
}

func GetFoodAndSideDishDetails(foodNames, sideDishNames []string) ([]models.Food, []models.SideDish, error) {
	var foods []models.Food
	var sideDishes []models.SideDish

	for _, foodName := range foodNames {
		foodDetails, err := grpc.FoodServiceClient.GetFoodDetailsByName(context.Background(), &food_proto.FoodIdRequest{FoodName: foodName})
		if err != nil {
			return nil, nil, err
		}
		food := models.Food{
			Name:         foodDetails.Name,
			Description:  foodDetails.Description,
			CategoryName: foodDetails.Category,
			MealName:     foodDetails.Meal,
		}
		food.ID = uint(foodDetails.Id)
		foods = append(foods, food)
	}

	for _, sideDishName := range sideDishNames {
		sideDishDetails, err := grpc.FoodServiceClient.GetSideDishDetailsByName(context.Background(), &food_proto.SideDishIdRequest{SideDishName: sideDishName})
		if err != nil {
			return nil, nil, err
		}
		sideDish := models.SideDish{
			Name:        sideDishDetails.Name,
			Description: sideDishDetails.Description,
		}
		sideDish.ID = uint(sideDishDetails.Id)
		sideDishes = append(sideDishes, sideDish)
	}

	return foods, sideDishes, nil
}

func ConvertUserDataToUser(userData *user_proto.GetUserInfoResponse) models.User {
	user := models.User{
		Username:     userData.UserData.Username,
		Email:        userData.UserData.Email,
		PhoneNumber:  userData.UserData.PhoneNumber,
		NationalCode: userData.UserData.NationalCode,
	}
	user.ID = uint(userData.UserData.UserId)
	return user
}

func GetUsernameFromJWT(cookie string) (string, error) {
	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return "", fmt.Errorf("failed to parse JWT token: %v", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", fmt.Errorf("invalid JWT token")
	}

	username, ok := claims["username"].(string)
	if !ok {
		return "", fmt.Errorf("username not found in JWT claims")
	}

	return username, nil
}

func AuthenticateUser(cookie string) (*user_proto.AuthenticateUserResponse, error) {
	if cookie == "" {
		return nil, fmt.Errorf("JWT cookie not found")
	}

	if err := grpc.InitializeUserGRPCClient(); err != nil {
		return nil, fmt.Errorf("failed to initialize gRPC client: %v", err)
	}

	return grpc.UserServiceClient.AuthenticateUser(context.Background(), &user_proto.AuthenticateUserRequest{JwtToken: cookie})
}
