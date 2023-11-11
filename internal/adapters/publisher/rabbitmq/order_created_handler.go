package rabbitmq

import (
	"context"
	"github.com/igoramorim/go-practice-clean-arch/pkg/ddd"
	"github.com/streadway/amqp"
	"sync"
)

func NewOrderCreatedHandler(rabbitMQChannel *amqp.Channel) *OrderCreatedHandler {
	return &OrderCreatedHandler{rabbitMQChannel: rabbitMQChannel}
}

type OrderCreatedHandler struct {
	rabbitMQChannel *amqp.Channel
}

func (h *OrderCreatedHandler) Handle(ctx context.Context, event ddd.Event, wg *sync.WaitGroup) error {
	defer wg.Done()

	msgRabbitmq := amqp.Publishing{
		ContentType: "application/json",
		Body:        event.Payload(),
	}

	return h.rabbitMQChannel.Publish("amq.direct", "", false, false, msgRabbitmq)
}
