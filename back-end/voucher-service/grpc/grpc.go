package grpc

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/nimbolism/software-restaurant/back-end/database"
	voucher_proto "github.com/nimbolism/software-restaurant/back-end/voucher-service/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	voucher_proto.UnimplementedVoucherServiceServer
}

var log = logrus.New() // Initialize the logger

// StoreOrderDetails implements VoucherServiceServer.StoreOrderDetails
func (s *Server) StoreOrderDetails(ctx context.Context, req *voucher_proto.StoreOrderDetailsRequestHelper) (*voucher_proto.StoreOrderDetailsResponseHelper, error) {
	order := req.GetOrder()
	fmt.Println(order)
	// Marshal order to bytes
	orderJson, err := json.Marshal(order)
	if err != nil {
		return nil, status.Error(
			codes.NotFound, fmt.Sprintf("failed to convert order to json: %v", err),
		)
	}

	orderString := string(orderJson)
	fmt.Println(orderString)
	// Set order in Redis with 24-hour expiration
	redisClient := database.GetRedisClient()
	err = redisClient.Set(ctx, fmt.Sprintf("order:%d", order.Id), orderString, 24*time.Hour).Err()
	if err != nil {
		return nil, status.Error(
			codes.NotFound, fmt.Sprintf("failed to store order in Redis: %v", err),
		)
	}

	return &voucher_proto.StoreOrderDetailsResponseHelper{Success: true}, nil
}

func (s *Server) GetAllOrders(ctx context.Context, req *voucher_proto.GetAllOrdersRequestHelper) (*voucher_proto.GetAllOrdersResponseHelper, error) {
	// Get Redis client
	redisClient := database.GetRedisClient()

	// Get all order keys from Redis
	orderKeys, err := redisClient.Keys(ctx, "order:*").Result()
	if err != nil {
		log.WithFields(logrus.Fields{
			"error": err,
		}).Error("failed to get order keys from Redis")
		return nil, status.Error(
			codes.NotFound, fmt.Sprintf("failed to get order keys from Redis: %v", err),
		)
	}

	// Retrieve orders from Redis
	orders := make([]*voucher_proto.OrderHelper, len(orderKeys))
	var wg sync.WaitGroup
	errors := make(chan error)

	for i, key := range orderKeys {
		wg.Add(1)
		go func(i int, key string) {
			defer wg.Done()
			orderBytes, err := redisClient.Get(ctx, key).Bytes()
			if err != nil {
				errors <- fmt.Errorf("failed to get order from Redis: %v", err)
				return
			}

			order := &voucher_proto.OrderHelper{}
			err = json.Unmarshal(orderBytes, order)
			if err != nil {
				errors <- fmt.Errorf("failed to unmarshal order from Redis: %v", err)
				return
			}

			orders[i] = order
		}(i, key)
	}

	go func() {
		wg.Wait()
		close(errors)
	}()

	for err := range errors {
		log.WithFields(logrus.Fields{
			"error": err,
		}).Error("error in goroutine")
	}

	return &voucher_proto.GetAllOrdersResponseHelper{Orders: orders}, nil
}

func StartServer() {
	println("Starting gRPC server...")
	lis, err := net.Listen("tcp", ":50050")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	voucher_proto.RegisterVoucherServiceServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
