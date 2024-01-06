package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/NayronFerreira/cleanArq_challenge/pkg/events"
	amqp "github.com/rabbitmq/amqp091-go"
)

type DeleteOrderHandler struct {
	RabbitMQChannel *amqp.Channel
}

func NewDeleteOrderHandler(rabbitMQChannel *amqp.Channel) *DeleteOrderHandler {
	return &DeleteOrderHandler{
		RabbitMQChannel: rabbitMQChannel,
	}
}

func (h *DeleteOrderHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("DeleteOrder used: %v", event.GetPayload())
	jsonOutput, _ := json.Marshal(event.GetPayload())

	msgRabbitmq := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonOutput,
	}

	h.RabbitMQChannel.PublishWithContext(
		context.Background(),
		"amq.direct",
		"delete_order",
		false,
		false,
		msgRabbitmq,
	)
}
