package dorder

import (
	"context"
	"time"
)

type Repository interface {
	Create(ctx context.Context, order *Order) error
	FindAllByPage(ctx context.Context, page, limit int, sort string) ([]*Order, int64, error)
	Count(ctx context.Context) (int64, error)
}

// TODO: Unit tests.

func FromRepository(id int64, price, tax, finalPrice float64, createdAt time.Time) *Order {
	return &Order{
		id:         idFromRepository(id),
		price:      price,
		tax:        tax,
		finalPrice: finalPrice,
		createdAt:  createdAt,
	}
}

func idFromRepository(id int64) ID {
	return ID{value: id}
}
