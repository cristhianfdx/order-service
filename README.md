# 🧾 Order Service

A RESTful Order Service built in Go using the Gin framework. It supports order creation, persistence in PostgreSQL, asynchronous processing via RabbitMQ, and a Backoffice interface for visualization.

---

## ⚙️ Tech Stack

- 🐹 **Go** with [Gin](https://github.com/gin-gonic/gin) for HTTP handling
- 🐘 **PostgreSQL** for persistent order storage
- 📦 **RabbitMQ** for asynchronous event publishing
- 🧱 **GORM** as the ORM layer for database operations

---

## 🚀 Features

- Create and retrieve orders via API
- Orders are stored in a PostgreSQL database
- Each order is published to a **RabbitMQ queue**

---

## 📦 Endpoints

| Method | Route            | Description                     |
|--------|------------------|--------------------------------|
| POST   | `/orders`        | Create a new order             |
| GET    | `/orders/:id`    | Get an order by its ID         |

---

## 🧪 CURL Examples

### ✅ Create a New Order

```bash
curl -X POST http://localhost:8080/orders \
  -H "Content-Type: application/json" \
  -d '{
    "amount": 50000,
    "item": "TV"
  }'
