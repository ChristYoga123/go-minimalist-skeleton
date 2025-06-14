# Simple Go Skeleton

A lightweight and clean Go project skeleton that provides a solid foundation for building microservices. This skeleton is designed to be minimal yet extensible, focusing on essential components without unnecessary dependencies.

## Features

- 🚀 Built with [Go Fiber](https://gofiber.io/) for high-performance HTTP routing
- 📦 [GORM](https://gorm.io/) integration for database operations
- 🏗️ Clean architecture folder structure
- 🔍 Health check endpoint included
- 🛠️ Minimal dependencies for microservice development
- 📝 Automatic database migrations with GORM

## Project Structure

```
.
├── app/                    # Application core
│   ├── entities/          # Domain models/entities
│   ├── http/             # HTTP layer (controllers, middlewares)
│   ├── repositories/     # Data access layer
│   └── services/        # Business logic layer
├── configs/              # Configuration files
│   └── migration.go     # Database migration configuration
├── routes/              # Route definitions
├── utils/               # Utility functions
├── main.go             # Application entry point
└── go.mod              # Go module file
```

## Getting Started

1. Clone the repository
2. Install dependencies:
   ```bash
   go mod download
   ```
3. Set up your environment variables:
   ```bash
   export APP_PORT=8080
   export DB_HOST=localhost
   export DB_PORT=3306
   export DB_USER=root
   export DB_PASSWORD=password
   export DB_NAME=your_database
   ```
4. Run the application:
   ```bash
   go run main.go
   ```

## API Endpoints

- `GET /api/health` - Health check endpoint

## Architecture

This skeleton follows a clean architecture approach with the following layers:

- **Entities**: Core business objects
- **Repositories**: Data access layer
- **Services**: Business logic implementation
- **HTTP Controllers**: Request handling and response formatting

## Database Migrations

The project uses GORM's auto-migration feature. To add new models to the migration:

1. Create your model in `app/entities/`
2. Register your model in `configs/migration.go`:
   ```go
   func RunMigration(db *gorm.DB) {
       err := db.AutoMigrate(
           &YourModel{},
           // Add more models here
       )
       // ...
   }
   ```

The migrations will run automatically when the application starts.

## Why This Skeleton?

- **Minimal Dependencies**: No unnecessary packages or features
- **Microservice Ready**: Perfect for building microservices
- **Clean Structure**: Clear separation of concerns
- **Easy to Extend**: Add your own models and features as needed

## Adding New Features

1. Create your entity (models/dtos) in `app/entities/`
2. Register your models in `configs/migration.go` for database migration
3. Add repository in `app/repositories/`
4. Implement business logic in `app/services/`
5. Create controller in `app/http/controllers/`
6. Add routes in `routes/api.go`