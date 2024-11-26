package postgres

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/yourusername/task-queue/internal/queue"
)

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage(db *sql.DB) *PostgresStorage {
	return &PostgresStorage{db: db}
}

func (p *PostgresStorage) CreateTask(ctx context.Context, task *queue.Task) error {
	query := `
        INSERT INTO tasks (id, name, payload, priority, status, retries, max_retries, created_at, delay_until)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
    `

	_, err := p.db.ExecContext(ctx, query,
		task.ID, task.Name, task.Payload, task.Priority,
		task.Status, task.Retries, task.MaxRetries,
		task.CreatedAt, task.DelayUntil)

	return err
}
