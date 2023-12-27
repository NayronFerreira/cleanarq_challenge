package web

import (
	"encoding/json"
	"net/http"

	"github.com/NayronFerreira/cleanArq_challenge/internal/entity"
	"github.com/NayronFerreira/cleanArq_challenge/internal/usecase"
)

type WebListOrdersHandler struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewWebListOrdersHandler(
	OrderRepository entity.OrderRepositoryInterface,
) *WebListOrdersHandler {
	return &WebListOrdersHandler{
		OrderRepository: OrderRepository,
	}
}

func (h *WebListOrdersHandler) List(res http.ResponseWriter, req *http.Request) {
	orderUseCase := usecase.NewListOrderUseCase(h.OrderRepository)
	output, err := orderUseCase.Execute()
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(res).Encode(output)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}
