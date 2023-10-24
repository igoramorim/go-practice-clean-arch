package dorder

import (
	"github.com/igoramorim/go-practice-clean-arch/pkg/ddd"
	"time"
)

func New(id int64, price, tax float64) (*Order, error) {
	orderID, err := NewID(id)
	if err != nil {
		return nil, err
	}

	if price <= 0 {
		return nil, ErrInvalidPrice
	}

	if tax <= 0 {
		return nil, ErrInvalidTax
	}

	order := &Order{
		id:         orderID,
		price:      price,
		tax:        tax,
		finalPrice: 0,
		createdAt:  time.Now().UTC(),
	}

	orderCreatedEvent := NewOrderCreatedEvent(order)
	order.defaultAggregate.AddEvent(orderCreatedEvent)

	return order, nil
}

type Order struct {
	id               ID
	price            float64
	tax              float64
	finalPrice       float64
	createdAt        time.Time
	defaultAggregate ddd.DefaultAggregate
}

func (o *Order) ID() int64 {
	return o.id.value
}

func (o *Order) Price() float64 {
	return o.price
}

func (o *Order) Tax() float64 {
	return o.tax
}

func (o *Order) FinalPrice() float64 {
	return o.finalPrice
}

func (o *Order) CreatedAt() time.Time {
	return o.createdAt
}

func (o *Order) Events() []ddd.Event {
	return o.defaultAggregate.Events()
}

func (o *Order) CalculateFinalPrice() error {
	finalPrice := o.price + o.tax
	if finalPrice <= 0 {
		return ErrInvalidFinalPrice
	}

	o.finalPrice = finalPrice
	return nil
}
