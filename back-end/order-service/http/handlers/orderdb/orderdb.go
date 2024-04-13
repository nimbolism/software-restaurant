package orderdb

import (
	"context"
	"encoding/json"
	"fmt"

	card_proto "github.com/nimbolism/software-restaurant/back-end/card-service/proto"
	"github.com/nimbolism/software-restaurant/back-end/database/models"
	"github.com/nimbolism/software-restaurant/back-end/gutils/postgresapp"
	"github.com/nimbolism/software-restaurant/back-end/order-service/grpc"
	"github.com/nimbolism/software-restaurant/back-end/order-service/http/handlers/utils"
	user_proto "github.com/nimbolism/software-restaurant/back-end/user-service/proto"
	"gorm.io/gorm"
)

func OrderHandler(msgBody []byte) error {
	var requestData struct {
		Paid   bool   `json:"paid"`
		Cookie string `json:"cookie"`
		Body   []byte `json:"body"`
	}
	if err := json.Unmarshal(msgBody, &requestData); err != nil {
		return fmt.Errorf("failed to decode message body: %v", err)
	}

	// Parse request body
	var bodyData struct {
		FoodNames     []string `json:"food_names"`
		SideDishNames []string `json:"side_dish_names"`
	}
	if err := json.Unmarshal(requestData.Body, &bodyData); err != nil {
		return fmt.Errorf("failed to decode body data: %v", err)
	}

	// Initialize gRPC clients
	if err := utils.InitializeGRPCClients(); err != nil {
		return fmt.Errorf("failed to initialize gRPC clients: %v", err)
	}

	// Get user information using gRPC call
	userInformationResponse, err := grpc.UserServiceClient.GetUserInfo(context.Background(), &user_proto.GetUserInfoRequest{JwtToken: requestData.Cookie})
	if err != nil {
		return fmt.Errorf("could not get user data: %v", err)
	}
	user := utils.ConvertUserDataToUser(userInformationResponse)

	// Get card details
	cardDetails, err := utils.GetCardDetails(requestData.Cookie)
	if err != nil {
		return fmt.Errorf("could not get card data: %v", err)
	}

	// Check if user is blacklisted
	if cardDetails.BlackListed {
		return fmt.Errorf("you are blacklisted and cannot place an order")
	}

	// Update card reserves if not paid and card is verified
	if !requestData.Paid && !cardDetails.Verified {
		return fmt.Errorf("you are not verified and cannot place order without paying")
	}

	// Get food and side dish details
	foods, sideDishes, err := utils.GetFoodAndSideDishDetails(bodyData.FoodNames, bodyData.SideDishNames)
	if err != nil {
		return err
	}

	if len(foods) == 0 {
		return fmt.Errorf("no food provided")
	}

	// Create order object
	order := models.Order{
		Foods:      foods,
		SideDishes: sideDishes,
		User:       user,
		UserID:     user.ID,
		Paid:       requestData.Paid,
	}

	// Save order to the database
	db := postgresapp.DB
	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&order).Error; err != nil {
			return err
		}

		if !requestData.Paid && cardDetails.Verified {
			_, err := grpc.CardServiceClient.UpdateReserves(context.Background(), &card_proto.UpdateReservesRequest{JwtToken: requestData.Cookie, ReservesChange: 1})
			if err != nil {
				// Rollback the order creation
				tx.Delete(&order)
				return fmt.Errorf("could not place order: %v", err)
			}
		}

		return nil
	}); err != nil {
		return err
	}

	if err := utils.SaveOrder(&order, db); err != nil {
		return fmt.Errorf("could not save order in redis db: %v", err)
	}

	return nil
}
