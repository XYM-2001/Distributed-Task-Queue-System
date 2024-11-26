package queue

import (
	"context"
	"time"

	"go.uber.org/zap"
)

type Worker struct {
	ID            string
	Tasks         chan *Task
	Quit          chan bool
	Logger        *zap.Logger
	MessageBroker MessageBroker
	Storage       TaskStorage
}

func NewWorker(id string, broker MessageBroker, storage TaskStorage, logger *zap.Logger) *Worker {
	return &Worker{
		ID:            id,
		Tasks:         make(chan *Task),
		Quit:          make(chan bool),
		Logger:        logger,
		MessageBroker: broker,
		Storage:       storage,
	}
}

func (w *Worker) Start(ctx context.Context) {
	w.Logger.Info("Starting worker", zap.String("worker_id", w.ID))

	for {
		select {
		case task := <-w.Tasks:
			w.processTask(ctx, task)
		case <-w.Quit:
			w.Logger.Info("Worker stopping", zap.String("worker_id", w.ID))
			return
		case <-ctx.Done():
			w.Logger.Info("Context cancelled, worker stopping", zap.String("worker_id", w.ID))
			return
		}
	}
}

func (w *Worker) processTask(ctx context.Context, task *Task) {
	w.Logger.Info("Processing task",
		zap.String("task_id", task.ID),
		zap.String("worker_id", w.ID))

	task.Status = TaskStatusRunning
	w.Storage.UpdateTask(ctx, task)

	// Simulate task processing
	time.Sleep(2 * time.Second)

	// Task processing logic would go here
	success := true // Replace with actual task execution

	if success {
		task.Status = TaskStatusCompleted
	} else {
		task.Retries++
		if task.Retries >= task.MaxRetries {
			task.Status = TaskStatusFailed
		} else {
			task.Status = TaskStatusPending
			task.DelayUntil = time.Now().Add(time.Second * time.Duration(task.Retries*5))
		}
	}

	w.Storage.UpdateTask(ctx, task)
}
