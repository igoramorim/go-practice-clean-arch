package ddd

import (
	"context"
	"github.com/pkg/errors"
	"sync"
)

// TODO: Add unit tests.

type EventDispatcher interface {
	Register(event Event, handlers ...EventHandler) error
	Dispatch(ctx context.Context, aggregate Aggregate) error
}

type EventHandler interface {
	Handle(ctx context.Context, event Event, wg *sync.WaitGroup) error
}

type DefaultEventDispatcher struct {
	handlers map[string][]EventHandler
}

func NewDefaultEventDispatcher() *DefaultEventDispatcher {
	return &DefaultEventDispatcher{
		handlers: make(map[string][]EventHandler),
	}
}

var ErrHandlerAlreadyRegistered = errors.New("handler already registered")

func (ed *DefaultEventDispatcher) Register(event Event, handlers ...EventHandler) error {
	for _, handler := range handlers {
		if _, ok := ed.handlers[event.Name()]; ok {
			for _, h := range ed.handlers[event.Name()] {
				if h == handler {
					return ErrHandlerAlreadyRegistered
				}
			}
		}
		ed.handlers[event.Name()] = append(ed.handlers[event.Name()], handler)
	}
	return nil
}

func (ed *DefaultEventDispatcher) Dispatch(ctx context.Context, aggregate Aggregate) error {
	events := aggregate.Events()
	for _, event := range events {
		if handlers, ok := ed.handlers[event.Name()]; ok {
			wg := &sync.WaitGroup{}
			for _, handler := range handlers {
				wg.Add(1)
				go handler.Handle(ctx, event, wg) // TODO: Handle error
			}
			wg.Wait()
		}
	}
	return nil
}
