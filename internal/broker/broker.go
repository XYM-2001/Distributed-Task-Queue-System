package broker

import (
	"context"

	"github.com/yourusername/task-queue/internal/queue"
)

type MessageBroker interface {
	PublishTask(ctx context.Context, task *queue.Task) error
	ConsumeTask(ctx context.Context) (<-chan *queue.Task, error)
	Close() error
}
