package web

import (
	"encoding/json"
	"net/http"

	"github.com/NayronFerreira/cleanArq_challenge/internal/entity"
	"github.com/NayronFerreira/cleanArq_challenge/internal/usecase"
	"github.com/NayronFerreira/cleanArq_challenge/pkg/events"
)

type WebDeleteOrderHandler struct {
	EventDispatcher  events.EventDispatcherInterface
	OrderRepository  entity.OrderRepositoryInterface
	DeleteOrderEvent events.EventInterface
}

func NewWebDeleteOrderHandler(
	EventDispatcher events.EventDispatcherInterface,
	OrderRepository entity.OrderRepositoryInterface,
	DeleteOrderEvent events.EventInterface,
) *WebDeleteOrderHandler {
	return &WebDeleteOrderHandler{
		EventDispatcher:  EventDispatcher,
		OrderRepository:  OrderRepository,
		DeleteOrderEvent: DeleteOrderEvent,
	}
}

func (h *WebDeleteOrderHandler) DeleteOrder(res http.ResponseWriter, req *http.Request) {
	var dto usecase.OrderInputDTO
	if err := json.NewDecoder(req.Body).Decode(&dto); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	orderUseCase := usecase.NewDeleteOrderUseCase(h.OrderRepository, h.DeleteOrderEvent, h.EventDispatcher)
	if err := orderUseCase.Execute(dto); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}
