package rabbitmq

import (
	"encoding/json"
	"log"

	"github.com/cristhianfdx/order-service/internal/domain"
	"github.com/cristhianfdx/order-service/internal/ports"
	"github.com/streadway/amqp"
)

type Publisher struct {
	channel *amqp.Channel
	queue   amqp.Queue
}

func NewPublisher(conn *amqp.Connection) ports.OrderPublisher {
	ch, _ := conn.Channel()
	q, _ := ch.QueueDeclare("order_created", false, false, false, false, nil)
	return &Publisher{channel: ch, queue: q}
}

func (p *Publisher) PublishOrderCreated(order *domain.Order) error {
	body, _ := json.Marshal(order)
	err := p.channel.Publish(
		"", p.queue.Name, false, false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})

	if err != nil {
		log.Println("Error publishing: ", err)
	}
	return err
}
