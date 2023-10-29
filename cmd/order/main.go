package main

import (
	"context"
	"github.com/igoramorim/go-practice-clean-arch/internal/adapters/repository/mysqlorder"
	"github.com/igoramorim/go-practice-clean-arch/internal/adapters/rest/restorder"
	"github.com/igoramorim/go-practice-clean-arch/internal/adapters/rest/webserver"
	"github.com/igoramorim/go-practice-clean-arch/internal/application"
	"github.com/igoramorim/go-practice-clean-arch/pkg/ddd"
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
	if err := webServer.Run(); err != nil {
		panic(err)
	}
}

type mockPublisher struct {
}

func (mock *mockPublisher) Publish(ctx context.Context, aggregates ...ddd.Aggregate) error {
	return nil
}
