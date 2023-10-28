package dorder

import "context"

type CreateOrderUseCaseInput struct {
	ID    int64
	Price float64
	Tax   float64
}

type CreateOrderUseCaseOutput struct {
	ID         int64
	Price      float64
	Tax        float64
	FinalPrice float64
	CreatedAt  string
}

type CreateOrderUseCase interface {
	Execute(ctx context.Context, input CreateOrderUseCaseInput) (CreateOrderUseCaseOutput, error)
}
