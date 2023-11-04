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
	Paging Paging
	Orders []FindAllOrdersByPageUseCaseOutputItem
}

type Paging struct {
	Limit  int64
	Offset int64
	Total  int64
}

type FindAllOrdersByPageUseCaseOutputItem struct {
	ID         int64
	Price      float64
	Tax        float64
	FinalPrice float64
	CreatedAt  string
}

func (out *FindAllOrdersByPageUseCaseOutputItem) Map(order *Order) {
	out.ID = order.id.value
	out.Price = order.price
	out.Tax = order.tax
	out.FinalPrice = order.finalPrice
	out.CreatedAt = order.createdAt.Format(time.RFC3339Nano)
}

type FindAllOrdersByPageUseCase interface {
	Execute(ctx context.Context, input FindAllOrdersByPageUseCaseInput) (FindAllOrdersByPageUseCaseOutput, error)
}
