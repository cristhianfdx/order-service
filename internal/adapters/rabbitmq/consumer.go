package rabbitmq

import (
	"encoding/json"
	"log"

	"github.com/cristhianfdx/order-service/internal/ports"
	"github.com/streadway/amqp"
)

func StartConsumer(conn *amqp.Connection, service ports.OrderService) {
	ch, _ := conn.Channel()
	q, _ := ch.QueueDeclare("order_created", false, false, false, false, nil)
	msgs, _ := ch.Consume(q.Name, "", true, false, false, false, nil)

	go func() {
		for d := range msgs {
			var payload struct {
				ID string `json:"id"`
			}
			if err := json.Unmarshal(d.Body, &payload); err != nil {
				log.Println("Error decoding message", err)
				continue
			}
			log.Println("Received order created event: ", payload.ID)
			service.MarkOrderComplete(payload.ID)
		}
	}()
}
