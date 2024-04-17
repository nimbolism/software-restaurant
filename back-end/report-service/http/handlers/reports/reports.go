package reports

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/nimbolism/software-restaurant/back-end/gutils"
	order_proto "github.com/nimbolism/software-restaurant/back-end/order-service/proto"
	"github.com/nimbolism/software-restaurant/back-end/report-service/grpc"
	"github.com/nimbolism/software-restaurant/back-end/report-service/http/handlers/utils"
	user_proto "github.com/nimbolism/software-restaurant/back-end/user-service/proto"
	voucher_proto "github.com/nimbolism/software-restaurant/back-end/voucher-service/proto"
)

type DateRange struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

func GetAllUsers(c *fiber.Ctx) error {
	cookie := gutils.GetCookie(c, "jwt")
	accessLevel, err := utils.GetAccessLevel(cookie)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{"error": fmt.Sprintf("failed to get access level information: %v", err)})
	}
	if accessLevel < 2 {
		return c.Status(fiber.StatusUnauthorized).JSON(
			fiber.Map{"error": "you are not autherized"})
	}
	if err := grpc.InitializeUserGRPCClient(); err != nil {
		return fmt.Errorf("failed to initialize user gRPC client: %w", err)
	}
	users, err := grpc.UserServiceClient.GetAllUsers(context.Background(), &user_proto.GetAllUsersRequest{JwtToken: cookie})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{"error": fmt.Sprintf("failed to get users information: %v", err)})
	}
	return c.JSON(users)
}
func GetAllOrders(c *fiber.Ctx) error {
	cookie := gutils.GetCookie(c, "jwt")
	accessLevel, err := utils.GetAccessLevel(cookie)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{"error": fmt.Sprintf("failed to get access level information: %v", err)})
	}
	if accessLevel < 2 {
		return c.Status(fiber.StatusUnauthorized).JSON(
			fiber.Map{"error": "you are not autherized"})
	}
	if err := grpc.InitializeOrderGRPCClient(); err != nil {
		return fmt.Errorf("failed to initialize order gRPC client: %w", err)
	}
	orders, err := grpc.OrderServiceClient.GetAllOrders(context.Background(), &order_proto.GetAllOrdersRequest{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{"error": fmt.Sprintf("failed to get orders information: %v", err)})
	}
	return c.JSON(orders)
}
func GetAllOrdersByTime(c *fiber.Ctx) error {
	cookie := gutils.GetCookie(c, "jwt")
	accessLevel, err := utils.GetAccessLevel(cookie)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{"error": fmt.Sprintf("failed to get access level information: %v", err)})
	}
	if accessLevel < 2 {
		return c.Status(fiber.StatusUnauthorized).JSON(
			fiber.Map{"error": "you are not autherized"})
	}
	// Parse request body
	var dateRange DateRange
	if err := c.BodyParser(&dateRange); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{"error": "invalid request body"})
	}
	startTime, err := time.Parse(time.RFC3339, dateRange.Start)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{"error": "invalid start date format"})
	}
	endTime, err := time.Parse(time.RFC3339, dateRange.End)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{"error": "invalid end date format"})
	}
	if err := grpc.InitializeOrderGRPCClient(); err != nil {
		return fmt.Errorf("failed to initialize order gRPC client: %w", err)
	}
	orders, err := grpc.OrderServiceClient.GetAllOrdersBetweenTimestamps(context.Background(), &order_proto.GetAllOrdersBetweenTimestampsRequest{
		StartTime: &timestamp.Timestamp{Seconds: startTime.Unix()},
		EndTime:   &timestamp.Timestamp{Seconds: endTime.Unix()},
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{"error": fmt.Sprintf("failed to get orders information: %v", err)})
	}
	return c.JSON(orders)
}
func GetAllOrdersVoucher(c *fiber.Ctx) error {
	cookie := gutils.GetCookie(c, "jwt")
	accessLevel, err := utils.GetAccessLevel(cookie)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{"error": fmt.Sprintf("failed to get access level information: %v", err)})
	}
	if accessLevel < 2 {
		return c.Status(fiber.StatusUnauthorized).JSON(
			fiber.Map{"error": "you are not autherized"})
	}
	if err := grpc.InitializeVoucherGRPCClient(); err != nil {
		return fmt.Errorf("failed to initialize order gRPC client: %w", err)
	}
	orders, err := grpc.VoucherServiceClient.GetAllOrders(context.Background(), &voucher_proto.GetAllOrdersRequestHelper{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{"error": fmt.Sprintf("failed to get orders information: %v", err)})
	}

	return c.JSON(orders)
}
