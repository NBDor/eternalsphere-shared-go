# Eternal Sphere Shared Go Library

Core shared libraries for Eternal Sphere microservices.

## Structure

```
eternalsphere-shared-go/
├── auth/          # Authentication utilities and middleware
├── database/      # Database connection handlers
├── events/        # Event bus interfaces and models
├── models/        # Shared data models
├── errors/        # Common error types
├── config/        # Configuration management
└── utils/         # General utilities
```

## Setup

```bash
go mod init github.com/yourusername/eternalsphere-shared-go
go mod tidy
```

## Modules

### Auth
- JWT middleware for service authentication
- Token generation and validation
- Role-based access control

### Database
- PostgreSQL connection management
- Redis caching interface
- MongoDB logging interface

### Events
- Event bus protocols
- Message models for inter-service communication

### Models
- Shared data structures
- Cross-service type definitions

## Usage

Import in other services:

```go
import "github.com/NBDor/eternalsphere-shared-go/auth/middleware"
```

## Testing

```bash
go test ./...
```