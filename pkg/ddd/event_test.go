package ddd_test

import (
	"github.com/igoramorim/go-practice-clean-arch/pkg/ddd"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testEvent struct{}

func (e testEvent) Name() string {
	return "test_event_name"
}

func TestDefaultAggregate_Events(t *testing.T) {
	event := testEvent{}
	defaultAggregate := ddd.DefaultAggregate{}
	defaultAggregate.AddEvent(event)

	events := defaultAggregate.Events()
	assert.Len(t, events, 1)
	assert.Empty(t, defaultAggregate.Events())
}
