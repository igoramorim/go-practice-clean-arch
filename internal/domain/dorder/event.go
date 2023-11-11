package dorder

import (
	"encoding/json"
	"time"
)

type OrderCreatedEvent struct {
	ID         int64   `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
	CreatedAt  string  `json:"created_at"`
}

func (e OrderCreatedEvent) Name() string {
	return "order_created"
}

func (e OrderCreatedEvent) Payload() []byte {
	payload, _ := json.Marshal(e)
	return payload
}

func NewOrderCreatedEvent(order *Order) OrderCreatedEvent {
	return OrderCreatedEvent{
		ID:         order.id.value,
		Price:      order.price,
		Tax:        order.tax,
		FinalPrice: order.finalPrice,
		CreatedAt:  order.createdAt.UTC().Format(time.RFC3339Nano),
	}
}
