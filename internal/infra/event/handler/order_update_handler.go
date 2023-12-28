package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/NayronFerreira/cleanArq_challenge/pkg/events"
	amqp "github.com/rabbitmq/amqp091-go"
)

type OrderUpdateHandler struct {
	RabbitMQChannel *amqp.Channel
}

func NewOrderUpdateHandler(rabbitMQChannel *amqp.Channel) *OrderUpdateHandler {
	return &OrderUpdateHandler{
		RabbitMQChannel: rabbitMQChannel,
	}
}

func (h *OrderUpdateHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Order updated: %v", event.GetPayload())
	jsonOutput, _ := json.Marshal(event.GetPayload())

	msgRabbitmq := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonOutput,
	}

	h.RabbitMQChannel.PublishWithContext(
		context.Background(),
		"amq.direct",
		"update_order",
		false,
		false,
		msgRabbitmq,
	)
}
