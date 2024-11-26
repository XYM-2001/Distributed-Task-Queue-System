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

func (r *RedisStorage) UpdateTask(ctx context.Context, task *queue.Task) error {
	return r.CreateTask(ctx, task)
}

func (r *RedisStorage) DeleteTask(ctx context.Context, id string) error {
	key := "task:" + id
	return r.client.Del(ctx, key).Err()
}

func (r *RedisStorage) ListTasks(ctx context.Context, limit, offset int) ([]*queue.Task, error) {
	pattern := "task:*"
	keys, err := r.client.Keys(ctx, pattern).Result()
	if err != nil {
		return nil, err
	}

	if offset >= len(keys) {
		return []*queue.Task{}, nil
	}

	end := offset + limit
	if end > len(keys) {
		end = len(keys)
	}

	tasks := make([]*queue.Task, 0, limit)
	for _, key := range keys[offset:end] {
		task, err := r.GetTask(ctx, key[5:])
		if err != nil {
			continue
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}
