package redis

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
	"github.com/yourusername/task-queue/internal/queue"
)

type RedisStorage struct {
	client *redis.Client
}

func NewRedisStorage(client *redis.Client) *RedisStorage {
	return &RedisStorage{client: client}
}

func (r *RedisStorage) CreateTask(ctx context.Context, task *queue.Task) error {
	data, err := json.Marshal(task)
	if err != nil {
		return err
	}

	key := "task:" + task.ID
	return r.client.Set(ctx, key, data, 0).Err()
}

func (r *RedisStorage) GetTask(ctx context.Context, id string) (*queue.Task, error) {
	key := "task:" + id
	data, err := r.client.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}

	var task queue.Task
	if err := json.Unmarshal(data, &task); err != nil {
		return nil, err
	}

	return &task, nil
}
