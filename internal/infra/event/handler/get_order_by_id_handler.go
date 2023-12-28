package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/NayronFerreira/cleanArq_challenge/pkg/events"
	amqp "github.com/rabbitmq/amqp091-go"
)

type GetOrderByIDHandler struct {
	RabbitMQChannel *amqp.Channel
}

func NewGetOrderByIDHandler(rabbitMQChannel *amqp.Channel) *GetOrderByIDHandler {
	return &GetOrderByIDHandler{
		RabbitMQChannel: rabbitMQChannel,
	}
}

func (h *GetOrderByIDHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("GetByID used: %v", event.GetPayload())
	jsonOutput, _ := json.Marshal(event.GetPayload())

	msgRabbitmq := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonOutput,
	}

	h.RabbitMQChannel.PublishWithContext(
		context.Background(),
		"amq.direct",
		"get_by_id",
		false,
		false,
		msgRabbitmq,
	)
}
