package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/yourusername/task-queue/internal/broker/rabbitmq"
	"github.com/yourusername/task-queue/internal/config"
	"github.com/yourusername/task-queue/internal/queue"
	"github.com/yourusername/task-queue/internal/storage/redis"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	cfg, err := config.Load("configs/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	// Initialize storage
	redisClient := redis.NewClient(cfg.Redis)
	storage := redis.NewRedisStorage(redisClient)

	// Initialize message broker
	broker, err := rabbitmq.NewRabbitMQ(cfg.RabbitMQ.URL())
	if err != nil {
		log.Fatal(err)
	}
	defer broker.Close()

	// Create worker
	worker := queue.NewWorker(
		os.Getenv("WORKER_ID"),
		broker,
		storage,
		logger,
	)

	// Setup context with cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle shutdown gracefully
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		cancel()
	}()

	// Start worker
	worker.Start(ctx)
}
