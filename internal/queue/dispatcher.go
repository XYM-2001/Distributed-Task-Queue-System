package queue

import (
	"context"
	"fmt"
	"sync"

	"go.uber.org/zap"
)

type Dispatcher struct {
	workers   []*Worker
	taskQueue chan *Task
	mutex     sync.RWMutex
	logger    *zap.Logger
	broker    MessageBroker
	storage   TaskStorage
}

func NewDispatcher(numWorkers int, broker MessageBroker, storage TaskStorage, logger *zap.Logger) *Dispatcher {
	d := &Dispatcher{
		workers:   make([]*Worker, numWorkers),
		taskQueue: make(chan *Task, 100),
		logger:    logger,
		broker:    broker,
		storage:   storage,
	}

	for i := 0; i < numWorkers; i++ {
		worker := NewWorker(fmt.Sprintf("worker-%d", i), broker, storage, logger)
		d.workers[i] = worker
	}

	return d
}

func (d *Dispatcher) Start(ctx context.Context) {
	d.logger.Info("Starting dispatcher")

	// Start all workers
	for _, worker := range d.workers {
		go worker.Start(ctx)
	}

	// Start task distribution
	go d.distribute(ctx)
}

func (d *Dispatcher) distribute(ctx context.Context) {
	for {
		select {
		case task := <-d.taskQueue:
			// Find available worker
			worker := d.getAvailableWorker()
			if worker != nil {
				worker.Tasks <- task
			}
		case <-ctx.Done():
			d.logger.Info("Dispatcher stopping")
			return
		}
	}
}

func (d *Dispatcher) getAvailableWorker() *Worker {
	d.mutex.RLock()
	defer d.mutex.RUnlock()

	// Simple round-robin selection
	// In a production environment, you might want to implement more sophisticated selection
	for _, worker := range d.workers {
		select {
		case worker.Tasks <- struct{}{}:
			return worker
		default:
			continue
		}
	}
	return nil
}

func (d *Dispatcher) EnqueueTask(task *Task) error {
	d.taskQueue <- task
	return nil
}
