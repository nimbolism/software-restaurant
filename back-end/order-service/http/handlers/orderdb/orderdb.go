package orderdb

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	card_proto "github.com/nimbolism/software-restaurant/back-end/card-service/proto"
	"github.com/nimbolism/software-restaurant/back-end/database/models"
	"github.com/nimbolism/software-restaurant/back-end/gutils"
	"github.com/nimbolism/software-restaurant/back-end/gutils/postgresapp"
	"github.com/nimbolism/software-restaurant/back-end/order-service/grpc"
	"github.com/nimbolism/software-restaurant/back-end/order-service/http/handlers/utils"
	user_proto "github.com/nimbolism/software-restaurant/back-end/user-service/proto"
	"gorm.io/gorm"
)

func OrderHandler(msgBody []byte) (string, error) {
	var requestData struct {
		Paid   bool   `json:"paid"`
		Cookie string `json:"cookie"`
		Body   []byte `json:"body"`
	}
	if err := json.Unmarshal(msgBody, &requestData); err != nil {
		return "", fmt.Errorf("failed to decode message body: %v", err)
	}
	username, err := utils.GetUsernameFromJWT(requestData.Cookie)
	if err != nil {
		return username, fmt.Errorf("failed to get username from body data: %v", err)
	}
	// Parse request body
	var bodyData struct {
		FoodNames     []string `json:"food_names"`
		SideDishNames []string `json:"side_dish_names"`
	}
	if err := json.Unmarshal(requestData.Body, &bodyData); err != nil {
		return username, fmt.Errorf("failed to decode body data: %v", err)
	}

	// Initialize gRPC clients
	if err := utils.InitializeGRPCClients(); err != nil {
		return username, fmt.Errorf("failed to initialize gRPC clients: %v", err)
	}

	// Get user information using gRPC call
	userInformationResponse, err := grpc.UserServiceClient.GetUserInfo(context.Background(), &user_proto.GetUserInfoRequest{JwtToken: requestData.Cookie})
	if err != nil {
		return username, fmt.Errorf("could not get user data: %v", err)
	}
	user := utils.ConvertUserDataToUser(userInformationResponse)

	// Get card details
	cardDetails, err := utils.GetCardDetails(requestData.Cookie)
	if err != nil {
		return username, fmt.Errorf("could not get card data: %v", err)
	}

	// Check if user is blacklisted
	if cardDetails.BlackListed {
		return username, fmt.Errorf("you are blacklisted and cannot place an order")
	}

	// Update card reserves if not paid and card is verified
	if !requestData.Paid && !cardDetails.Verified {
		return username, fmt.Errorf("you are not verified and cannot place order without paying")
	}

	// Get food and side dish details
	foods, sideDishes, err := utils.GetFoodAndSideDishDetails(bodyData.FoodNames, bodyData.SideDishNames)
	if err != nil {
		return username, err
	}

	if len(foods) == 0 {
		return username, fmt.Errorf("no food provided")
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
			if err := utils.SaveOrder(&order, db); err != nil {
				tx.Delete(&order)
				return fmt.Errorf("could not save order in redis db: %v", err)
			}
		}

		return nil
	}); err != nil {
		return username, err
	}

	return username, nil
}

func GetOrders(c *fiber.Ctx) error {
	if err := grpc.InitializeUserGRPCClient(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": fmt.Sprintf("start the connection: %v", err)})
	}
	cookie := gutils.GetCookie(c, "jwt")
	authRespose, err := grpc.UserServiceClient.AuthenticateUser(context.Background(), &user_proto.AuthenticateUserRequest{JwtToken: cookie})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": fmt.Sprintf("could not get the username: %v", err)})
	}
	userID := authRespose.UserId
	var orders []models.Order
	db := postgresapp.DB
	if err := db.Preload("Foods").Preload("SideDishes").Find(&orders, "user_id = ?", userID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": fmt.Sprintf("could not get any orders: %v", err)})
	}

	var orderInfo []fiber.Map
	for _, order := range orders {
		orderData := fiber.Map{
			"paid":        order.Paid,
			"created_at":  order.CreatedAt,
			"foods":       order.Foods,
			"side_dishes": order.SideDishes,
		}
		orderInfo = append(orderInfo, orderData)
	}

	return c.JSON(orderInfo)
}

func GetFailedOrders(c *fiber.Ctx) error {
	cookie := gutils.GetCookie(c, "jwt")
	username, err := utils.GetUsernameFromJWT(cookie)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": fmt.Sprintf("could not get username: %v", err)})
	}
	var orderFails []models.OrderFail
	db := postgresapp.DB
	if err := db.Where("username = ?", username).Find(&orderFails).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": fmt.Sprintf("could not get failed orders: %v", err)})
	}

	return c.JSON(orderFails)
}
