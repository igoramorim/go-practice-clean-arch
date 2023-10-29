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

type FindAllOrdersByPageUseCaseInput struct {
	Page  int
	Limit int
	Sort  string
}

type FindAllOrdersByPageUseCaseOutput struct {
	ID         int64
	Price      float64
	Tax        float64
	FinalPrice float64
	CreatedAt  string
}

func (out *FindAllOrdersByPageUseCaseOutput) Map(order *Order) {
	out.ID = order.id.value
	out.Price = order.price
	out.Tax = order.tax
	out.FinalPrice = order.finalPrice
	out.CreatedAt = order.createdAt.Format(time.RFC3339Nano)
}

type FindAllOrdersByPageUseCase interface {
	Execute(ctx context.Context, input FindAllOrdersByPageUseCaseInput) ([]FindAllOrdersByPageUseCaseOutput, error)
}
