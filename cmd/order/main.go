package main

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/igoramorim/go-practice-clean-arch/config"
	"github.com/igoramorim/go-practice-clean-arch/internal/adapters/graph"
	"github.com/igoramorim/go-practice-clean-arch/internal/adapters/grpc/grpcorder"
	"github.com/igoramorim/go-practice-clean-arch/internal/adapters/grpc/pb"
	"github.com/igoramorim/go-practice-clean-arch/internal/adapters/publisher/rabbitmq"
	"github.com/igoramorim/go-practice-clean-arch/internal/adapters/repository/mysql"
	"github.com/igoramorim/go-practice-clean-arch/internal/adapters/repository/mysqlorder"
	"github.com/igoramorim/go-practice-clean-arch/internal/adapters/rest/restorder"
	"github.com/igoramorim/go-practice-clean-arch/internal/adapters/rest/webserver"
	"github.com/igoramorim/go-practice-clean-arch/internal/application"
	"github.com/igoramorim/go-practice-clean-arch/internal/domain/dorder"
	"github.com/igoramorim/go-practice-clean-arch/pkg/ddd"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
)

func main() {
	// TODO: Add basic logging.

	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	// DB Connections
	db := mysql.OpenConn(cfg.DBMySQLUser, cfg.DBMySQLPass, cfg.DBMySQLHost, cfg.DBMySQLPort, cfg.DBMySQLDatabase)
	defer db.Close()

	// Publishers
	rabbitMQChannel := rabbitmq.OpenChannel(cfg.RabbitMQUser, cfg.RabbitMQPass, cfg.RabbitMQHost, cfg.RabbitMQPort)
	defer rabbitMQChannel.Close()

	eventDispatcher := ddd.NewDefaultEventDispatcher()
	err = eventDispatcher.Register(dorder.OrderCreatedEvent{}, rabbitmq.NewOrderCreatedHandler(rabbitMQChannel))
	if err != nil {
		panic(err)
	}

	// Repositories
	orderRepository := mysqlorder.New(db)

	// UseCases
	createOrderUseCase := application.NewCreateOrderService(orderRepository, eventDispatcher)
	findAllOrdersByPageUseCase := application.NewFindAllOrdersByPageService(orderRepository)

	// Web Handlers
	webOrderHandler := restorder.NewHandler(createOrderUseCase, findAllOrdersByPageUseCase)

	// Web Server
	webServer := webserver.New(cfg.WebServerPort)
	webServer.AddHandler(http.MethodPost, "/orders", webOrderHandler.CreateOrder)
	webServer.AddHandler(http.MethodGet, "/orders", webOrderHandler.FindAllByPage)

	// GraphQL
	graphQLServer := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CreateOrderUseCase:         createOrderUseCase,
		FindAllOrdersByPageUseCase: findAllOrdersByPageUseCase,
	}}))
	webServer.AddHandler(http.MethodPost, "/query", func(w http.ResponseWriter, r *http.Request) {
		graphQLServer.ServeHTTP(w, r)
	})
	webServer.AddHandler(http.MethodGet, "/playground", playground.Handler("GraphQL Playground", "/query"))

	// Web Server Up!
	go func() {
		if err = webServer.Run(); err != nil {
			panic(err)
		}
	}()

	// Grpc
	grpcServer := grpc.NewServer()
	grpcOrderService := grpcorder.NewService(createOrderUseCase, findAllOrdersByPageUseCase)
	pb.RegisterOrderServiceServer(grpcServer, grpcOrderService)
	reflection.Register(grpcServer) // For evans
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.GrpcServerPort))
	if err != nil {
		panic(err)
	}
	if err = grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
