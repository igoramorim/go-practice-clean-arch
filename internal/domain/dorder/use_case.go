package dorder

import (
	"context"
	"time"
)

type CreateOrderUseCaseInput struct {
	ID    int64   `json:"id"`
	Price float64 `json:"price"`
	Tax   float64 `json:"tax"`
}

type CreateOrderUseCaseOutput struct {
	ID         int64   `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
	CreatedAt  string  `json:"created_at"`
}

type CreateOrderUseCase interface {
	Execute(ctx context.Context, input CreateOrderUseCaseInput) (CreateOrderUseCaseOutput, error)
}

type FindAllOrdersByPageUseCaseInput struct {
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
	Sort  string `json:"sort"`
}

type FindAllOrdersByPageUseCaseOutput struct {
	Paging Paging                                 `json:"paging"`
	Orders []FindAllOrdersByPageUseCaseOutputItem `json:"orders"`
}

type Paging struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
	Total  int64 `json:"total"`
}

type FindAllOrdersByPageUseCaseOutputItem struct {
	ID         int64   `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
	CreatedAt  string  `json:"created_at"`
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
