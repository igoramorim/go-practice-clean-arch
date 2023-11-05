package model

import "github.com/igoramorim/go-practice-clean-arch/internal/domain/dorder"

// TODO: Add unit test.

func MapCreateOrderOutput(out dorder.CreateOrderUseCaseOutput) *Order {
	return &Order{
		ID:         int(out.ID),
		Price:      out.Price,
		Tax:        out.Tax,
		FinalPrice: out.FinalPrice,
		CreatedAt:  out.CreatedAt,
	}
}

func MapFindAllOrdersByPageOutput(out dorder.FindAllOrdersByPageUseCaseOutput) *FindAllOrdersByPageOutput {
	orders := make([]*Order, 0, len(out.Orders))
	for _, o := range out.Orders {
		orders = append(orders, &Order{
			ID:         int(o.ID),
			Price:      o.Price,
			Tax:        o.Tax,
			FinalPrice: o.FinalPrice,
			CreatedAt:  o.CreatedAt,
		})
	}

	return &FindAllOrdersByPageOutput{
		Paging: &Paging{
			Limit:  int(out.Paging.Limit),
			Offset: int(out.Paging.Offset),
			Total:  int(out.Paging.Total),
		},
		Orders: orders,
	}
}
