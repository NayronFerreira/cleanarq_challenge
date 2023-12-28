package usecase

import (
	"github.com/NayronFerreira/cleanArq_challenge/internal/entity"
	"github.com/NayronFerreira/cleanArq_challenge/pkg/events"
)

type UpdateOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	OrderUpdate     events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewUpdateOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	OrderUpdate events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *UpdateOrderUseCase {
	return &UpdateOrderUseCase{
		OrderRepository: OrderRepository,
		OrderUpdate:     OrderUpdate,
		EventDispatcher: EventDispatcher,
	}
}

func (c *UpdateOrderUseCase) Execute(input OrderInputDTO) (OrderOutputDTO, error) {

	order := entity.Order{
		ID:    input.ID,
		Price: input.Price,
		Tax:   input.Tax,
	}
	order.CalculateFinalPrice()
	orderUpdate, err := c.OrderRepository.UpdateOrder(&order)
	if err != nil {
		return OrderOutputDTO{}, err
	}

	dto := OrderOutputDTO{
		ID:         orderUpdate.ID,
		Price:      orderUpdate.Price,
		Tax:        orderUpdate.Tax,
		FinalPrice: orderUpdate.FinalPrice,
	}

	c.OrderUpdate.SetPayload(dto)
	c.EventDispatcher.Dispatch(c.OrderUpdate)

	return dto, nil
}
