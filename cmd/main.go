package main

import (
	"log"

	postgresAdapter "github.com/cristhianfdx/order-service/internal/adapters/postgres"
	rabbitmqAdapter "github.com/cristhianfdx/order-service/internal/adapters/rabbitmq"
	api "github.com/cristhianfdx/order-service/internal/api"
	app "github.com/cristhianfdx/order-service/internal/app"

	gin "github.com/gin-gonic/gin"
	amqp "github.com/streadway/amqp"
	postgresDriver "gorm.io/driver/postgres"
	gorm "gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=orders port=5432 sslmode=disable"
	db, err := gorm.Open(postgresDriver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connevt to database: %v", err)
	}

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	orderRepo := postgresAdapter.NewOrderRepository(db)
	orderPub := rabbitmqAdapter.NewPublisher(conn)

	service := app.NewOrderService(orderRepo, orderPub)
	handler := api.NewOrderHandler(service)

	go rabbitmqAdapter.StartConsumer(conn, service)

	r := gin.Default()
	r.POST("/orders", handler.CreateOrder)
	r.GET("/orders/:id", handler.GetOrder)

	r.Run(":8080")
}
