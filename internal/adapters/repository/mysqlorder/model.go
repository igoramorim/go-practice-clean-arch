package mysqlorder

import (
	"github.com/igoramorim/go-practice-clean-arch/internal/domain/dorder"
	"time"
)

type PersistableOrder struct {
	ID         int64
	Price      float64
	Tax        float64
	FinalPrice float64
	CreatedAt  time.Time
}

// TODO: Unit tests.

func PersistableFromDomain(order *dorder.Order) PersistableOrder {
	return PersistableOrder{
		ID:         order.ID(),
		Price:      order.Price(),
		Tax:        order.Tax(),
		FinalPrice: order.Tax(),
		CreatedAt:  order.CreatedAt(),
	}
}

// TODO: Unit tests.

func PersistableToDomain(persistable PersistableOrder) *dorder.Order {
	return dorder.FromRepository(persistable.ID, persistable.Price, persistable.Tax, persistable.FinalPrice,
		persistable.CreatedAt)
}
