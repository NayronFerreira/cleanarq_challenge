package usecase

import (
	"github.com/NayronFerreira/cleanArq_challenge/internal/entity"
	"github.com/NayronFerreira/cleanArq_challenge/pkg/events"
)

type ListOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	OrderListed     events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewListOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	OrderListed events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *ListOrderUseCase {
	return &ListOrderUseCase{
		OrderRepository: OrderRepository,
		OrderListed:     OrderListed,
		EventDispatcher: EventDispatcher,
	}
}

func (c *ListOrderUseCase) Execute() ([]OrderOutputDTO, error) {
	orders, err := c.OrderRepository.GetOrders()
	if err != nil {
		return nil, err
	}

	var output []OrderOutputDTO
	for _, order := range orders {
		output = append(output, OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}

	c.OrderListed.SetPayload(output)
	c.EventDispatcher.Dispatch(c.OrderListed)

	return output, nil
}
