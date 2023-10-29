package dorder

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, order *Order) error
	FindAllByPage(ctx context.Context, page, limit int, sort string) ([]*Order, int64, error)
	Count(ctx context.Context) (int64, error)
}
