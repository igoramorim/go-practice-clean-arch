package dorder

import "context"

type CreateOrderUseCaseInput struct {
	ID    string
	Price float64
	Tax   float64
}

type CreateOrderUseCaseOutput struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

type CreateOrderUseCase interface {
	Execute(ctx context.Context, input CreateOrderUseCaseInput) (CreateOrderUseCaseOutput, error)
}
