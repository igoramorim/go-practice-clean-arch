package dorder

import "context"

type Repository interface {
	Save(ctx context.Context, order *Order) error
	FindAllByPage(ctx context.Context, page, limit int, sort string) ([]*Order, error)
}
