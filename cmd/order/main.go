package main

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/igoramorim/go-practice-clean-arch/internal/adapters/graph"
	"github.com/igoramorim/go-practice-clean-arch/internal/adapters/grpc/grpcorder"
	"github.com/igoramorim/go-practice-clean-arch/internal/adapters/grpc/pb"
	"github.com/igoramorim/go-practice-clean-arch/internal/adapters/repository/mysqlorder"
	"github.com/igoramorim/go-practice-clean-arch/internal/adapters/rest/restorder"
	"github.com/igoramorim/go-practice-clean-arch/internal/adapters/rest/webserver"
	"github.com/igoramorim/go-practice-clean-arch/internal/application"
	"github.com/igoramorim/go-practice-clean-arch/pkg/ddd"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
)

func main() {
	// TODO: Load configs

	// TODO: SQl conn

	// TODO: Publishers

	// Repositories
	orderRepository := mysqlorder.New(nil)

	// UseCases
	createOrderUseCase := application.NewCreateOrderService(orderRepository, &mockPublisher{})
	findAllOrdersByPageUseCase := application.NewFindAllOrdersByPageService(orderRepository)

	// Web Handlers
	webOrderHandler := restorder.NewHandler(createOrderUseCase, findAllOrdersByPageUseCase)

	// Web Server
	webServer := webserver.New("8080")
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
		if err := webServer.Run(); err != nil {
			panic(err)
		}
	}()

	// Grpc
	grpcServer := grpc.NewServer()
	grpcOrderService := grpcorder.NewService(createOrderUseCase, findAllOrdersByPageUseCase)
	pb.RegisterOrderServiceServer(grpcServer, grpcOrderService)
	reflection.Register(grpcServer) // For evans
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", "50051"))
	if err != nil {
		panic(err)
	}
	if err = grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}

type mockPublisher struct {
}

func (mock *mockPublisher) Publish(ctx context.Context, aggregates ...ddd.Aggregate) error {
	return nil
}
