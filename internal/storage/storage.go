package storage

import (
	"context"

	"github.com/yourusername/task-queue/internal/queue"
)

type TaskStorage interface {
	CreateTask(ctx context.Context, task *queue.Task) error
	GetTask(ctx context.Context, id string) (*queue.Task, error)
	UpdateTask(ctx context.Context, task *queue.Task) error
	DeleteTask(ctx context.Context, id string) error
	ListTasks(ctx context.Context, limit, offset int) ([]*queue.Task, error)
}

type WorkerStorage interface {
	RegisterWorker(ctx context.Context, worker *queue.Worker) error
	UpdateWorkerStatus(ctx context.Context, workerID string, status string) error
	ListWorkers(ctx context.Context) ([]*queue.Worker, error)
}
