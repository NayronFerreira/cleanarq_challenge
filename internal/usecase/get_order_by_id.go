package usecase

import (
	"github.com/NayronFerreira/cleanArq_challenge/internal/entity"
	"github.com/NayronFerreira/cleanArq_challenge/pkg/events"
)

type GetOrderByIDUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	GetOrderByID    events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewGetOrderByIDUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	GetOrderByID events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *GetOrderByIDUseCase {
	return &GetOrderByIDUseCase{
		OrderRepository: OrderRepository,
		GetOrderByID:    GetOrderByID,
		EventDispatcher: EventDispatcher,
	}
}

func (c *GetOrderByIDUseCase) Execute(input OrderInputDTO) (OrderOutputDTO, error) {

	order, err := c.OrderRepository.GetOrderByID(input.ID)
	if err != nil {
		return OrderOutputDTO{}, err
	}

	dto := OrderOutputDTO{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}

	c.GetOrderByID.SetPayload(dto)
	c.EventDispatcher.Dispatch(c.GetOrderByID)

	return dto, nil
}
