# Distributed Task Queue System

A scalable and distributed task queue system built with Go, designed to handle background jobs and asynchronous tasks across multiple workers.

## Features

- **Task Scheduling**
  - Priority-based task queuing
  - Delayed task execution
  - Automatic retry mechanism
  
- **Worker Management**
  - Distributed worker pool
  - Automatic worker registration
  - Health monitoring
  
- **Storage Options**
  - Redis for fast, in-memory operations
  - PostgreSQL for persistent storage
  
- **Message Broker**
  - RabbitMQ integration for reliable message delivery
  - Support for distributed task distribution
  
- **Monitoring**
  - Real-time metrics with Prometheus
  - Web-based dashboard
  - Detailed logging

## Prerequisites

- Go 1.21+
- Docker and Docker Compose
- Redis
- PostgreSQL
- RabbitMQ

## Installation

1. Clone the repository:
