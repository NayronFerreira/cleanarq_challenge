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
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
	wire.Bind(new(events.EventInterface), new(*event.OrderList)),
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
