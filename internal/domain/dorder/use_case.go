package dorder

import (
	"context"
	"time"
)

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

type FindAllByPageUseCaseInput struct {
	Page  int
	Limit int
	Sort  string
}

type FindAllByPageUseCaseOutput struct {
	ID         int64
	Price      float64
	Tax        float64
	FinalPrice float64
	CreatedAt  string
}

func (out *FindAllByPageUseCaseOutput) Map(order *Order) {
	out.ID = order.id.value
	out.Price = order.price
	out.Tax = order.tax
	out.FinalPrice = order.finalPrice
	out.CreatedAt = order.createdAt.Format(time.RFC3339Nano)
}

type FindAllByPageUseCase interface {
	Execute(ctx context.Context, input FindAllByPageUseCaseInput) ([]FindAllByPageUseCaseOutput, error)
}
