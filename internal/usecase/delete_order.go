package usecase

import (
	"github.com/NayronFerreira/cleanArq_challenge/internal/entity"
	"github.com/NayronFerreira/cleanArq_challenge/pkg/events"
)

type DeleteOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	DeleteOrder     events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewDeleteOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	DeleteOrder events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *DeleteOrderUseCase {
	return &DeleteOrderUseCase{
		OrderRepository: OrderRepository,
		DeleteOrder:     DeleteOrder,
		EventDispatcher: EventDispatcher,
	}
}

func (c *DeleteOrderUseCase) Execute(input OrderInputDTO) error {

	if err := c.OrderRepository.DeleteOrder(input.ID); err != nil {
		return err
	}

	c.DeleteOrder.SetPayload("This order was deleted id: " + input.ID)
	c.EventDispatcher.Dispatch(c.DeleteOrder)

	return nil
}
