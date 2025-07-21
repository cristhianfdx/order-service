
# Order Service

A Go-based microservice for managing orders, using PostgreSQL for data persistence, RabbitMQ for message queuing, and Gin-Gonic as the HTTP server framework.

## Features

- RESTful API for managing orders.
- PostgreSQL with GORM for database operations.
- RabbitMQ integration for publishing order events.
- Docker Compose setup for easy local development.
- Secure connection with SSH keygen.
- Access to RabbitMQ management web interface.

## Technologies

- Go (Gin, GORM)
- PostgreSQL
- RabbitMQ
- Docker & Docker Compose

## Setup Instructions

### 1. Clone the repository

```bash
git clone git@github.com:your-username/order-service.git
cd order-service
```

### 2. Generate SSH key (if needed for GitHub access)

```bash
ssh-keygen -t ed25519 -C "your_email@example.com"
eval "$(ssh-agent -s)"
ssh-add ~/.ssh/id_ed25519
cat ~/.ssh/id_ed25519.pub
# Add this public key to your GitHub account
```

### 3. Run the services

```bash
docker-compose up --build
go run cmd/main.go
```

### 4. Access RabbitMQ Management

RabbitMQ management UI is available at [http://localhost:15672](http://localhost:15672)

- **Username**: guest
- **Password**: guest

## API Endpoints

### Create Order

```bash
curl -X POST http://localhost:8080/orders   -H "Content-Type: application/json"   -d '{
    "item": "Calculator",
    "amount": 20000
}'
```

### Get Order by ID

```bash
curl http://localhost:8080/orders/1
```

## Environment Variables

You can adjust the following values in `.env`:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=orders
RABBITMQ_URL=amqp://guest:guest@rabbitmq:5672/
```

## Project Structure

```
.
├── internal/
│   ├── api/            # HTTP handlers
│   ├── service/        # Business logic (ports)
│   ├── repository/     # Persistence layer (adapters)
│   └── domain/         # Core domain entities
├── docker-compose.yml
├── main.go
└── README.md
```

## License

This project is licensed under the MIT License.
