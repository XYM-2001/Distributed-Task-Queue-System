package integration_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/yourusername/task-queue/internal/queue"
)

func TestWorkerProcessTask(t *testing.T) {
	ctx := context.Background()
	storage := setupTestRedisStorage(t)
	broker := setupTestMessageBroker(t)
	logger := setupTestLogger(t)

	worker := queue.NewWorker("test-worker", broker, storage, logger)

	task := queue.NewTask("test_task", []byte(`{"test": true}`), queue.MediumPriority)

	// Start worker
	go worker.Start(ctx)

	// Send task
	worker.Tasks <- task

	// Wait for processing
	time.Sleep(time.Second * 3)

	// Check task status
	processedTask, err := storage.GetTask(ctx, task.ID)
	assert.NoError(t, err)
	assert.Equal(t, queue.TaskStatusCompleted, processedTask.Status)
}
