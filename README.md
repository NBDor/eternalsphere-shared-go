# Eternal Sphere Shared Go Library

Core shared libraries for Eternal Sphere microservices.

## Prerequisites

- Go 1.21+
- PostgreSQL instance for running database tests
- Redis for caching tests
- MongoDB for event logging tests

## Setup

```bash
go mod init github.com/yourusername/eternalsphere-shared-go
go mod tidy
```

## Project Structure

```
eternalsphere-shared-go/
├── auth/          # Authentication utilities and middleware
│   ├── middleware/    # JWT and auth middleware
│   └── models/       # Auth-related models
├── database/      # Database connection handlers
│   ├── postgres/     # PostgreSQL connections
│   ├── redis/        # Redis caching
│   └── mongodb/      # MongoDB logging
├── events/        # Event bus interfaces
├── models/        # Shared data models
├── errors/        # Common error types
├── config/        # Configuration management
└── utils/         # General utilities
```

## Testing

### Database Tests
Required environment:
```bash
# PostgreSQL
export PG_HOST=localhost
export PG_PORT=5432
export PG_USER=test
export PG_PASSWORD=test
export PG_DATABASE=test

# Run tests
go test ./...
```

## Usage Examples

### JWT Middleware
```go
import "github.com/yourusername/eternalsphere-shared-go/auth/middleware"

config := middleware.JWTConfig{
    SecretKey: "your-secret-key",
    TokenDuration: time.Hour * 24,
}
router.Use(middleware.JWTMiddleware(config))
```

### Database Connection
```go
import "github.com/yourusername/eternalsphere-shared-go/database/postgres"

config := postgres.Config{
    Host:     "localhost",
    Port:     5432,
    User:     "user",
    Password: "password",
    DBName:   "dbname",
    SSLMode:  "disable",
}

conn, err := postgres.NewConnection(config)
if err != nil {
    log.Fatal(err)
}
defer conn.Close()
```