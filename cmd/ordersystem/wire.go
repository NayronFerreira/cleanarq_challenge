//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/NayronFerreira/cleanArq_challenge/internal/entity"
	"github.com/NayronFerreira/cleanArq_challenge/internal/infra/database"
	"github.com/NayronFerreira/cleanArq_challenge/internal/infra/event"
	"github.com/NayronFerreira/cleanArq_challenge/internal/infra/web"
	"github.com/NayronFerreira/cleanArq_challenge/internal/usecase"
	"github.com/NayronFerreira/cleanArq_challenge/pkg/events"
	"github.com/google/wire"
)

var setOrderRepositoryDependency = wire.NewSet(
	database.NewOrderRepository,
	wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)),
)

var setEventDispatcherDependency = wire.NewSet(
	events.NewEventDispatcher,
	event.NewOrderCreated,
	event.NewOrderList,
	event.NewGetOrderByID,
	event.NewOrderUpdate,
	event.NewDeleteOrder,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
	wire.Bind(new(events.EventInterface), new(*event.OrderList)),
	wire.Bind(new(events.EventInterface), new(*event.GetOrderByID)),
	wire.Bind(new(events.EventInterface), new(*event.OrderUpdate)),
	wire.Bind(new(events.EventInterface), new(*event.DeleteOrder)),
	wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)),
)

var setOrderCreatedEvent = wire.NewSet(
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
)

var setOrderListedEvent = wire.NewSet(
	event.NewOrderList,
	wire.Bind(new(events.EventInterface), new(*event.OrderList)),
)

var setGetOrderByIDEvent = wire.NewSet(
	event.NewGetOrderByID,
	wire.Bind(new(events.EventInterface), new(*event.GetOrderByID)),
)

var setUpdateOrderEvent = wire.NewSet(
	event.NewOrderUpdate,
	wire.Bind(new(events.EventInterface), new(*event.OrderUpdate)),
)

var setDeleteOrderEvent = wire.NewSet(
	event.NewDeleteOrder,
	wire.Bind(new(events.EventInterface), new(*event.DeleteOrder)),
)

func NewCreateOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.CreateOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		usecase.NewCreateOrderUseCase,
	)
	return &usecase.CreateOrderUseCase{}
}

func NewListOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.ListOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderListedEvent,
		usecase.NewListOrderUseCase,
	)
	return &usecase.ListOrderUseCase{}
}

func NewGetOrderByIDUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.GetOrderByIDUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		setGetOrderByIDEvent,
		usecase.NewGetOrderByIDUseCase,
	)
	return &usecase.GetOrderByIDUseCase{}
}

func NewUpdateOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.UpdateOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		setUpdateOrderEvent,
		usecase.NewUpdateOrderUseCase,
	)
	return &usecase.UpdateOrderUseCase{}
}

func NewDeleteOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.DeleteOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		setDeleteOrderEvent,
		usecase.NewDeleteOrderUseCase,
	)
	return &usecase.DeleteOrderUseCase{}
}

func NewWebOrderHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebOrderHandler {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		web.NewWebOrderHandler,
	)
	return &web.WebOrderHandler{}
}

func NewWebListOrdersHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebListOrdersHandler {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderListedEvent,
		web.NewWebListOrdersHandler,
	)
	return &web.WebListOrdersHandler{}
}

func NewWebGetOrderByIDHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebGetOrderByIDHandler {
	wire.Build(
		setOrderRepositoryDependency,
		setGetOrderByIDEvent,
		web.NewWebGetOrderByIDHandler,
	)
	return &web.WebGetOrderByIDHandler{}
}

func NewWebUpdateOrderHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebUpdateOrderHandler {
	wire.Build(
		setOrderRepositoryDependency,
		setUpdateOrderEvent,
		web.NewWebUpdateOrderHandler,
	)
	return &web.WebUpdateOrderHandler{}
}

func NewWebDeleteOrderHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebDeleteOrderHandler {
	wire.Build(
		setOrderRepositoryDependency,
		setDeleteOrderEvent,
		web.NewWebDeleteOrderHandler,
	)
	return &web.WebDeleteOrderHandler{}
}
