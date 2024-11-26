package queue_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yourusername/task-queue/internal/queue"
)

func TestNewTask(t *testing.T) {
	name := "test_task"
	payload := []byte(`{"key": "value"}`)
	priority := queue.HighPriority

	task := queue.NewTask(name, payload, priority)

	assert.NotEmpty(t, task.ID)
	assert.Equal(t, name, task.Name)
	assert.Equal(t, payload, task.Payload)
	assert.Equal(t, priority, task.Priority)
	assert.Equal(t, queue.TaskStatusPending, task.Status)
	assert.Equal(t, 0, task.Retries)
	assert.Equal(t, 3, task.MaxRetries)
	assert.False(t, task.CreatedAt.IsZero())
}
