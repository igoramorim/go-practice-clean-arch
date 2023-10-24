package dorder_test

import (
	"errors"
	"github.com/igoramorim/go-practice-clean-arch/internal/domain/dorder"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewOrder(t *testing.T) {
	t.Run("Create an Order successfully", func(t *testing.T) {
		order, err := dorder.New(1, 10.99, 0.1)
		assert.NoError(t, err)
		assert.NotNil(t, order)
		assert.Equal(t, int64(1), order.ID())
		assert.Equal(t, 10.99, order.Price())
		assert.Equal(t, 0.1, order.Tax())
		assert.Equal(t, 0.0, order.FinalPrice())
		assert.NotEmpty(t, order.CreatedAt())

		events := order.Events()
		assert.Len(t, events, 1)
		assert.IsType(t, dorder.OrderCreatedEvent{}, events[0])
	})

	t.Run("Return error: invalid id", func(t *testing.T) {
		order, err := dorder.New(0, 10.99, 0.1)
		assert.True(t, errors.Is(err, dorder.ErrInvalidID))
		assert.Nil(t, order)
	})

	t.Run("Return error: invalid price", func(t *testing.T) {
		order, err := dorder.New(1, 0, 0.1)
		assert.True(t, errors.Is(err, dorder.ErrInvalidPrice))
		assert.Nil(t, order)
	})

	t.Run("Return error: invalid tax", func(t *testing.T) {
		order, err := dorder.New(1, 10.99, 0)
		assert.True(t, errors.Is(err, dorder.ErrInvalidTax))
		assert.Nil(t, order)
	})
}

func TestOrder_CalculateFinalPrice(t *testing.T) {
	t.Run("Calculate final price successfully", func(t *testing.T) {
		order, err := dorder.New(1, 10.99, 0.1)
		assert.NotNil(t, order)
		assert.NoError(t, err)

		err = order.CalculateFinalPrice()
		assert.NoError(t, err)
		assert.Equal(t, 11.09, order.FinalPrice())
	})
}
