package graph

import "github.com/igoramorim/go-practice-clean-arch/internal/domain/dorder"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CreateOrderUseCase         dorder.CreateOrderUseCase
	FindAllOrdersByPageUseCase dorder.FindAllOrdersByPageUseCase
}
