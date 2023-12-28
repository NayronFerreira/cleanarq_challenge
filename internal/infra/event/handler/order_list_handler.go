package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/NayronFerreira/cleanArq_challenge/pkg/events"
	amqp "github.com/rabbitmq/amqp091-go"
)

type OrderListHandler struct {
	RabbitMQChannel *amqp.Channel
}

func NewOrderListHandler(rabbitMQChannel *amqp.Channel) *OrderListHandler {
	return &OrderListHandler{
		RabbitMQChannel: rabbitMQChannel,
	}
}

func (h *OrderListHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Order Listed: %v", event.GetPayload())
	jsonOutput, _ := json.Marshal(event.GetPayload())

	msgRabbitmq := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonOutput,
	}

	h.RabbitMQChannel.PublishWithContext(
		context.Background(),
		"amq.direct",
		"list",
		false,
		false,
		msgRabbitmq,
	)
}
