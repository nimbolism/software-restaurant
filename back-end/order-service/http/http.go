package http

import (
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nimbolism/software-restaurant/back-end/database/models"
	"github.com/nimbolism/software-restaurant/back-end/gutils"
	"github.com/nimbolism/software-restaurant/back-end/gutils/postgresapp"
	"github.com/nimbolism/software-restaurant/back-end/order-service/http/handlers/orderdb"
	"github.com/nimbolism/software-restaurant/back-end/rabbitmq"
)

var rmq *rabbitmq.RabbitMQ

func init() {
	// Initialize RabbitMQ connection
	var err error
	rmq, err = rabbitmq.NewRabbitMQ()
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	// Declare a queue for order messages
	err = rmq.DeclareQueue("order_queue")
	if err != nil {
		log.Fatalf("Failed to declare queue: %v", err)
	}
}

func StartServer() {
	app := fiber.New()

	// /order for order service
	orderGroup := app.Group("/order")

	// root endpoint for testing
	orderGroup.Get("/", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/plain")
		return c.SendString("Welcome to the order Service!")
	})

	orderApis := orderGroup.Group("/api")
	orderApis.Post("/order/:paid", func(c *fiber.Ctx) error {
		// Extract the paid parameter from the URL
		paidString := c.Params("paid")
		paid, err := strconv.ParseBool(paidString)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid 'paid' parameter"})
		}
		cookie := gutils.GetCookie(c, "jwt")
		if cookie == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "JWT cookie not found"})
		}
		// Prepare the message payload including the paid parameter
		message := struct {
			Paid   bool   `json:"paid"`
			Cookie string `json:"cookie"`
			Body   []byte `json:"body"`
		}{
			Paid:   paid,
			Cookie: cookie,
			Body:   c.Body(),
		}

		// Marshal the message payload to JSON
		messageBody, err := json.Marshal(message)
		if err != nil {
			return err
		}

		// Publish the message payload to RabbitMQ
		err = rmq.Publish("order_queue", messageBody)
		if err != nil {
			return err
		}
		return c.SendString("Order submitted successfully!")
	})
	orderApis.Get("/order/failed", orderdb.GetFailedOrders)
	orderApis.Get("/order/success", orderdb.GetOrders)

	consumeMessages := func() {
		for {
			// Consume messages from the "order_queue" queue
			msgs, err := rmq.Consume("order_queue")
			if err != nil {
				log.Printf("Failed to consume messages from queue: %v", err)
				time.Sleep(5 * time.Second) // Wait for 5 seconds before retrying
				continue
			}
			for msg := range msgs {
				username, err := orderdb.OrderHandler(msg.Body)
				if err != nil {
					db := postgresapp.DB
					orderfail := models.OrderFail{
						Username: username,
						Error:    err.Error(),
					}
					db.Save(&orderfail)
				}
			}
		}
	}

	go consumeMessages()

	println("Starting order HTTP server...")
	if err := app.Listen(":8040"); err != nil {
		log.Fatalf("Failed to start order HTTP server: %v", err)
	}
}

func CloseRabbitMQ() {
	rmq.Close()
}
