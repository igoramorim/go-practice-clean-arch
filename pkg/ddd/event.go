package ddd

import "context"

type Event interface {
	Name() string
}

type Aggregate interface {
	Events() []Event
}

type DefaultAggregate struct {
	events []Event
}

func (da *DefaultAggregate) Events() []Event {
	events := da.events
	da.events = nil
	return events
}

func (da *DefaultAggregate) AddEvent(events ...Event) {
	da.events = append(da.events, events...)
}

type Publisher interface {
	Publish(ctx context.Context, aggregates ...Aggregate) error
}
