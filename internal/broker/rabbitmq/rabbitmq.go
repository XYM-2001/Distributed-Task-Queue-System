package rabbitmq

import (
	"context"
	"encoding/json"

	"github.com/streadway/amqp"
	"github.com/yourusername/task-queue/internal/queue"
)

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewRabbitMQ(url string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &RabbitMQ{
		conn:    conn,
		channel: ch,
	}, nil
}

func (r *RabbitMQ) PublishTask(ctx context.Context, task *queue.Task) error {
	data, err := json.Marshal(task)
	if err != nil {
		return err
	}

	return r.channel.Publish(
		"tasks",
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        data,
		},
	)
}
