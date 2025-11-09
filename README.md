# Sales Tracker

A REST API service for tracking sales data built with Go, featuring clean architecture, PostgreSQL database, and Docker containerization.

## Features

- **CRUD Operations**: Create, read, update, and delete sales records
- **Analytics**: Get statistical insights including sum, average, count, median, and 90th percentile
- **Clean Architecture**: Organized into handlers, services, and repositories layers
- **Database Migrations**: Automated schema management with Goose
- **Docker Support**: Containerized deployment with Docker Compose
- **CORS Enabled**: Configured for frontend integration (localhost:3000)
- **Validation**: Input validation using go-playground/validator
- **Logging**: Structured logging with zerolog

## Tech Stack

- **Language**: Go 1.25.3
- **Framework**: Gin (via wbf/ginext)
- **Database**: PostgreSQL
- **Migration Tool**: Goose
- **Containerization**: Docker & Docker Compose
- **Validation**: go-playground/validator

## API Endpoints

### Sales Management

- `POST /sales-tracker/api/items` - Create a new sale
- `GET /sales-tracker/api/items` - Get all sales
- `PUT /sales-tracker/api/items/:id` - Update a sale by ID
- `DELETE /sales-tracker/api/items/:id` - Delete a sale by ID

### Analytics

- `GET /sales-tracker/api/analytics` - Get sales analytics

## Database Schema

### Sale Table

```sql
CREATE TABLE sale (
    id SERIAL PRIMARY KEY,
    item VARCHAR(50) NOT NULL,
    income NUMERIC CHECK (income >= 0),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
```

**Indexes:**
- `idx_sale_created_at` on `created_at` column

**Triggers:**
- Auto-update `updated_at` on row updates

## Installation and Setup

1. **Clone the repository:**
   ```bash
   git clone https://github.com/avraam311/sales-tracker.git
   cd sales-tracker
   ```

2. **Environment Setup:**
   Create a `.env` file in the root directory:
   ```env

3. **Using Docker Compose (Recommended):**
   ```bash
   make up
   ```
   This will start the application, PostgreSQL database, and run migrations.

4. **Alternative: Local Development:**
   - Ensure PostgreSQL is running locally
   - Update `.env` with local database credentials
   - Run migrations manually
   - Start the application:
     ```bash
     go run cmd/main.go
     ```

## Usage

### API Request Examples

**Create a Sale:**
```bash
curl -X POST http://localhost:8080/sales-tracker/api/items \
  -H "Content-Type: application/json" \
  -d '{"item": "Product A", "income": 100.50}'
```

**Get All Sales:**
```bash
curl http://localhost:8080/sales-tracker/api/items
```

**Update a Sale:**
```bash
curl -X PUT http://localhost:8080/sales-tracker/api/items/1 \
  -H "Content-Type: application/json" \
  -d '{"item": "Updated Product", "income": 150.00}'
```

**Delete a Sale:**
```bash
curl -X DELETE http://localhost:8080/sales-tracker/api/items/1
```

**Get Analytics:**
```bash
curl http://localhost:8080/sales-tracker/api/analytics
```

## Development

### Available Commands

- `make lint` - Run Go vet and golangci-lint
- `make up` - Start services with Docker Compose
- `make buildup` - Build and start services
- `make down` - Stop and remove containers/volumes

### Project Structure

```
├── cmd/                    # Application entry point
│   ├── Dockerfile
│   └── main.go
├── config/                 # Configuration files
│   └── local.yaml
├── internal/
│   ├── api/                # HTTP layer
│   │   ├── handlers/       # Request handlers
│   │   ├── middlewares/    # HTTP middlewares
│   │   └── server/         # Server setup
│   ├── models/             # Data models
│   ├── repository/         # Data access
│   └── service/            # Business logic
├── migrations/             # Database migrations
├── docker-compose.yaml     # Docker services
├── go.mod                  # Go modules
├── go.sum                  # Go dependencies
├── Makefile                # Build commands
└── README.md
```

### Adding New Features

1. Define models in `internal/models/`
2. Implement repository methods in `internal/repository/`
3. Add business logic in `internal/service/`
4. Create HTTP handlers in `internal/api/handlers/`
5. Register routes in `internal/api/server/server.go`
6. Add database migrations if needed

## Docker

The application includes a multi-service Docker setup:

- **app**: Go application container
- **db**: PostgreSQL database
- **migrator**: Runs database migrations on startup

Services are orchestrated with health checks to ensure proper startup order.
