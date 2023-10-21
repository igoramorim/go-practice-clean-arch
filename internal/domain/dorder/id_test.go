package dorder_test

import (
	"errors"
	"github.com/igoramorim/go-practice-clean-arch/internal/domain/dorder"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewID(t *testing.T) {
	t.Run("Create an ID successfully", func(t *testing.T) {
		id, err := dorder.NewID(10)
		assert.Equal(t, int64(10), id.Value())
		assert.NoError(t, err)
	})

	t.Run("Return error creating ID with value 0", func(t *testing.T) {
		id, err := dorder.NewID(0)
		assert.True(t, id.IsEmpty())
		assert.True(t, errors.Is(err, dorder.ErrInvalidID))
	})
}

func TestID_IsEmpty(t *testing.T) {
	t.Run("Return true", func(t *testing.T) {
		id := dorder.IDEmptyValue()
		assert.True(t, id.IsEmpty())
	})

	t.Run("Return false", func(t *testing.T) {
		id, err := dorder.NewID(10)
		assert.False(t, id.IsEmpty())
		assert.NoError(t, err)
	})
}

func TestID_Equals(t *testing.T) {
	id, err := dorder.NewID(10)
	assert.NoError(t, err)

	otherIDSameValue, err := dorder.NewID(10)
	assert.NoError(t, err)

	otherID, err := dorder.NewID(20)
	assert.NoError(t, err)

	t.Run("Return true when same variables", func(t *testing.T) {
		assert.True(t, id.Equals(id))
	})

	t.Run("Return true when different variables but same value", func(t *testing.T) {
		assert.True(t, id.Equals(otherIDSameValue))
	})

	t.Run("Return false when different variables", func(t *testing.T) {
		assert.False(t, id.Equals(otherID))
	})
}

func TestIDEmptyValue(t *testing.T) {
	t.Run("Create an empty ID", func(t *testing.T) {
		id := dorder.IDEmptyValue()
		assert.True(t, id.IsEmpty())
	})
}
