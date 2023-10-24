package dorder_test

import (
	"github.com/igoramorim/go-practice-clean-arch/internal/domain/dorder"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrderCreatedEvent_Name(t *testing.T) {
	event := dorder.OrderCreatedEvent{}
	assert.Equal(t, "order_created", event.Name())
}

func TestNewOrderCreatedEvent(t *testing.T) {
	t.Run("Return an order created event successfully", func(t *testing.T) {
		order, err := dorder.New(1, 10.99, 0.1)
		assert.NoError(t, err)

		event := dorder.NewOrderCreatedEvent(order)
		assert.Equal(t, int64(1), event.ID)
		assert.Equal(t, 10.99, event.Price)
		assert.Equal(t, 0.1, event.Tax)
		assert.Equal(t, 0.0, event.FinalPrice)
		assert.NotEmpty(t, event.CreatedAt)
	})
}
