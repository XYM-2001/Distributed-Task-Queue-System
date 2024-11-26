package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/task-queue/internal/queue"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// Initialize your storage and message broker here
	storage := initializeStorage()
	broker := initializeMessageBroker()

	// Create dispatcher with 5 workers
	dispatcher := queue.NewDispatcher(5, broker, storage, logger)

	// Start dispatcher
	ctx := context.Background()
	go dispatcher.Start(ctx)

	// Initialize Gin router
	r := gin.Default()

	// API routes
	r.POST("/api/v1/tasks", func(c *gin.Context) {
		var taskRequest struct {
			Name     string          `json:"name"`
			Payload  json.RawMessage `json:"payload"`
			Priority queue.Priority  `json:"priority"`
		}

		if err := c.BindJSON(&taskRequest); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		task := queue.NewTask(taskRequest.Name, taskRequest.Payload, taskRequest.Priority)

		if err := dispatcher.EnqueueTask(task); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, task)
	})

	// Start server
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
