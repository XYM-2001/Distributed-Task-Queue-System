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

bash
git clone https://github.com/yourusername/task-queue.git
cd task-queue

2. Install dependencies:
bash
go mod download

3. Configure the application:
bash
cp configs/config.example.yaml configs/config.yaml

## Edit configs/config.yaml with your settings


## Running the Application

### Using Docker Compose

bash
docker compose up --build


### Manual Setup

1. Start the API server:

bash
go run cmd/api/main.go

2. Start the worker:

bash
go run cmd/worker/main.go

3. Start the dashboard:

bash
go run cmd/dashboard/main.go


## API Endpoints

### Tasks
- `POST /api/v1/tasks` - Create a new task
- `GET /api/v1/tasks/:id` - Get task status
- `GET /api/v1/tasks` - List tasks

### Workers
- `GET /api/v1/workers` - List active workers
- `GET /api/v1/workers/:id/status` - Get worker status

## Testing

Run the test suite:

bash
go test ./...

Run integration tests:

bash
go test ./tests/integration/...

## Monitoring

- Dashboard: http://localhost:3000
- Prometheus metrics: http://localhost:8080/metrics
- RabbitMQ management: http://localhost:15672

## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a new Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.


