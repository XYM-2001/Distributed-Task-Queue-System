package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	TasksProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "taskqueue_processed_tasks_total",
		Help: "The total number of processed tasks",
	})

	TasksFailedTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "taskqueue_failed_tasks_total",
		Help: "The total number of failed tasks",
	})

	TaskProcessingTime = promauto.NewHistogram(prometheus.HistogramOpts{
		Name:    "taskqueue_task_processing_seconds",
		Help:    "Time spent processing tasks",
		Buckets: prometheus.DefBuckets,
	})

	ActiveWorkers = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "taskqueue_active_workers",
		Help: "The number of currently active workers",
	})
)
