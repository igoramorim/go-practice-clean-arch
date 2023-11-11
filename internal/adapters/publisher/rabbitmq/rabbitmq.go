package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
)

func OpenChannel(user, pass, host, port string) *amqp.Channel {
	connString := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, pass, host, port)
	conn, err := amqp.Dial(connString)
	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	return ch
}
