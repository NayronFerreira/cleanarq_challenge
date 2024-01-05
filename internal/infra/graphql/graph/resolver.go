package graph

import "github.com/NayronFerreira/cleanArq_challenge/internal/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CreateOrderUseCase  usecase.CreateOrderUseCase
	ListOrderUseCase    usecase.ListOrderUseCase
	GetOrderByIDUseCase usecase.GetOrderByIDUseCase
	UpdateOrderUseCase  usecase.UpdateOrderUseCase
}
