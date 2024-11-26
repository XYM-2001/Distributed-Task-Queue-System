package queue

import (
	"time"

	"github.com/google/uuid"
)

type TaskStatus string

const (
	TaskStatusPending   TaskStatus = "pending"
	TaskStatusRunning   TaskStatus = "running"
	TaskStatusCompleted TaskStatus = "completed"
	TaskStatusFailed    TaskStatus = "failed"
)

type Priority int

const (
	LowPriority Priority = iota
	MediumPriority
	HighPriority
)

type Task struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	Payload    []byte     `json:"payload"`
	Priority   Priority   `json:"priority"`
	Status     TaskStatus `json:"status"`
	Retries    int        `json:"retries"`
	MaxRetries int        `json:"max_retries"`
	CreatedAt  time.Time  `json:"created_at"`
	DelayUntil time.Time  `json:"delay_until"`
	LastError  string     `json:"last_error,omitempty"`
}

func NewTask(name string, payload []byte, priority Priority) *Task {
	return &Task{
		ID:         uuid.New().String(),
		Name:       name,
		Payload:    payload,
		Priority:   priority,
		Status:     TaskStatusPending,
		MaxRetries: 3,
		CreatedAt:  time.Now(),
		DelayUntil: time.Now(),
	}
}

func (t *Task) CanProcess() bool {
	return t.Status != TaskStatusCompleted &&
		t.Retries < t.MaxRetries &&
		time.Now().After(t.DelayUntil)
}
