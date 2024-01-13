package main

import (
	"database/sql"
	"fmt"

	"net"
	"net/http"

	graphql_handler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/NayronFerreira/cleanArq_challenge/configs"
	"github.com/NayronFerreira/cleanArq_challenge/internal/infra/event/handler"

	"github.com/NayronFerreira/cleanArq_challenge/internal/infra/graphql/graph"
	"github.com/NayronFerreira/cleanArq_challenge/internal/infra/grpc/pb"
	"github.com/NayronFerreira/cleanArq_challenge/internal/infra/grpc/service"
	"github.com/NayronFerreira/cleanArq_challenge/internal/infra/web/webserver"
	"github.com/NayronFerreira/cleanArq_challenge/pkg/events"
	amqp "github.com/rabbitmq/amqp091-go"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rabbitMQChannel := getRabbitMQChannel()

	eventDispatcher := events.NewEventDispatcher()

	eventDispatcher.Register("OrderCreated", &handler.OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})
	eventDispatcher.Register("OrderList", &handler.OrderListHandler{
		RabbitMQChannel: rabbitMQChannel,
	})
	eventDispatcher.Register("GetOrderByID", &handler.GetOrderByIDHandler{
		RabbitMQChannel: rabbitMQChannel,
	})
	eventDispatcher.Register("OrderUpdate", &handler.OrderUpdateHandler{
		RabbitMQChannel: rabbitMQChannel,
	})
	eventDispatcher.Register("DeleteOrder", &handler.DeleteOrderHandler{
		RabbitMQChannel: rabbitMQChannel,
	})

	createOrderUseCase := NewCreateOrderUseCase(db, eventDispatcher)
	listOrdersUseCase := NewListOrderUseCase(db, eventDispatcher)
	getOrderByID := NewGetOrderByIDUseCase(db, eventDispatcher)
	updateOrderUseCase := NewUpdateOrderUseCase(db, eventDispatcher)
	deleteOrderUseCase := NewDeleteOrderUseCase(db, eventDispatcher)

	webserver := webserver.NewWebServer(configs.WebServerPort)

	webOrderHandler := NewWebOrderHandler(db, eventDispatcher)
	webserver.AddHandler("/order", webOrderHandler.Create, "POST")

	webListOrdersHandler := NewWebListOrdersHandler(db, eventDispatcher)
	webserver.AddHandler("/orders", webListOrdersHandler.List, "GET")

	webGetOrderByIDHandler := NewWebGetOrderByIDHandler(db, eventDispatcher)
	webserver.AddHandler("/orderbyid", webGetOrderByIDHandler.GetOrderByID, "GET")

	webOrderUpdateHandler := NewWebUpdateOrderHandler(db, eventDispatcher)
	webserver.AddHandler("/order/update", webOrderUpdateHandler.UpdateOrder, "POST")

	webDeleteUpdateHandler := NewWebDeleteOrderHandler(db, eventDispatcher)
	webserver.AddHandler("/order/delete", webDeleteUpdateHandler.DeleteOrder, "DELETE")

	fmt.Println("Starting web server on port", configs.WebServerPort)
	go webserver.Start()

	grpcServer := grpc.NewServer()

	orderService := service.NewOrderService(*createOrderUseCase, *listOrdersUseCase, *getOrderByID, *updateOrderUseCase, *deleteOrderUseCase)
	pb.RegisterOrderServiceServer(grpcServer, orderService)

	reflection.Register(grpcServer)

	fmt.Println("Starting gRPC server on port", configs.GRPCServerPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", configs.GRPCServerPort))
	if err != nil {
		panic(err)
	}
	go grpcServer.Serve(lis)

	srv := graphql_handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CreateOrderUseCase:  *createOrderUseCase,
		ListOrderUseCase:    *listOrdersUseCase,
		GetOrderByIDUseCase: *getOrderByID,
		UpdateOrderUseCase:  *updateOrderUseCase,
		DeleteOrderUseCase:  *deleteOrderUseCase,
	}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	fmt.Println("Starting GraphQL server on port", configs.GraphQLServerPort)
	http.ListenAndServe(":"+configs.GraphQLServerPort, nil)
}

func getRabbitMQChannel() *amqp.Channel {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch
}
