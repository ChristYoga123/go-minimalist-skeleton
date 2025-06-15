# Simple Go Skeleton

A lightweight and clean Go project skeleton that provides a solid foundation for building microservices. This skeleton is designed to be minimal yet extensible, focusing on essential components without unnecessary dependencies.

## Features

- 🚀 Built with [Go Fiber](https://gofiber.io/) for high-performance HTTP routing
- 📦 [GORM](https://gorm.io/) integration for database operations
- 🏗️ Simple and clean folder structure
- 🔍 Health check endpoint included
- 🛠️ Minimal dependencies for microservice development
- 📝 Automatic database migrations with GORM
- 🔄 Flexible database support (MySQL, PostgreSQL, SQLite, SQL Server)
- 🎯 Extensible for NoSQL databases (Redis, MongoDB, etc.)

## Project Structure

```
.
├── app/                    # Application core
│   ├── http/             # HTTP layer
│   │   ├── controllers/  # Controllers
│   │   │   ├── interfaces/  # Controller interfaces
│   │   │   └── healthControllerImpl.go
│   │   └── middlewares/  # HTTP middlewares
│   ├── repositories/     # Data access layer
│   │   ├── interfaces/   # Repository interfaces
│   │   └── healthRepositoryImpl.go
│   ├── services/        # Business logic layer
│   │   ├── interfaces/  # Service interfaces
│   │   └── healthServiceImpl.go
│   └── wire/           # Dependency injection
│       ├── wire.go     # Wire providers
│       └── wire_gen.go # Generated wire code
├── configs/              # Configuration files
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
   # Application
   export APP_PORT=8080
   
   # Database Configuration
   export DB_DRIVER=mysql     # Options: mysql, postgres, sqlite, sqlserver
   export DB_HOST=localhost
   export DB_PORT=3306
   export DB_DATABASE=your_database
   export DB_USERNAME=root
   export DB_PASSWORD=password
   ```

   Example configurations for different databases:
   ```bash
   # MySQL
   export DB_DRIVER=mysql
   export DB_PORT=3306

   # PostgreSQL
   export DB_DRIVER=postgres
   export DB_PORT=5432

   # SQLite
   export DB_DRIVER=sqlite
   export DB_DATABASE=gorm.db

   # SQL Server
   export DB_DRIVER=sqlserver
   export DB_PORT=1433
   ```

4. Run the application:
   ```bash
   go run main.go
   ```

## API Endpoints

- `GET /api/health` - Health check endpoint

## Architecture

This skeleton follows a clean architecture with dependency injection:

- **Wire**: Dependency injection using Google Wire
- **Controllers**: Request handling with interfaces
- **Services**: Business logic layer with interfaces
- **Repositories**: Data access layer with interfaces
- **Routes**: API endpoint definitions
- **Configs**: Application and database configuration
- **Utils**: Shared utility functions

The architecture implements:
- Repository Pattern for data access
- Service Layer for business logic
- Dependency Injection using Wire
- Interface-based design for better testing
- Clean separation of concerns

## Database Configuration

The project uses a factory pattern for database connections and implements the Singleton pattern for database instances. The configuration is handled through environment variables.

### Supported Databases
- MySQL
- PostgreSQL
- SQLite
- SQL Server

### Database Connection Pattern
The project implements:
- Factory pattern for database connections (`database.go`)
- Singleton pattern for database instances (`gorm.go`)
- Automatic migrations on startup
- Thread-safe initialization using `sync.Once`

### Adding NoSQL Support
The codebase is prepared for NoSQL integration with placeholders for:
- Redis
- MongoDB

To add a new database type:
1. Add the database driver case in `configs/database.go`
2. Implement the connection logic
3. Add the singleton instance in `configs/gorm.go`
4. Add the getter function in `configs/gorm.go`

Example of adding Redis support:

1. First, add Redis driver to your `go.mod`:
```bash
go get github.com/redis/go-redis/v9
```

2. Add Redis configuration to `configs/database.go`:
```go
// Add to DBConfig struct
type DBConfig struct {
    // ... existing fields ...
    RedisAddr     string `env:"REDIS_ADDR"`
    RedisPassword string `env:"REDIS_PASSWORD"`
    RedisDB       int    `env:"REDIS_DB"`
}

// Add to NewDatabaseConnection function
func NewDatabaseConnection(config *DBConfig) (db *gorm.DB, err error) {
    switch config.DBDriver {
    // ... existing cases ...
    case "redis":
        redisClient := redis.NewClient(&redis.Options{
            Addr:     config.RedisAddr,
            Password: config.RedisPassword,
            DB:       config.RedisDB,
        })
        // Test connection
        if err := redisClient.Ping(context.Background()).Err(); err != nil {
            return nil, fmt.Errorf("failed to connect to Redis: %w", err)
        }
        return nil, nil // Redis doesn't use GORM
    }
    // ... rest of the function
}
```

3. Add Redis singleton to `configs/gorm.go`:
```go
var (
    // ... existing variables ...
    redisClient *redis.Client
)

// Add to InitDatabase function
func InitDatabase() error {
    once.Do(func() {
        // ... existing code ...
        if cfg.DBDriver == "redis" {
            redisClient = redis.NewClient(&redis.Options{
                Addr:     cfg.RedisAddr,
                Password: cfg.RedisPassword,
                DB:       cfg.RedisDB,
            })
        }
    })
    return err
}

// Add Redis getter
func GetRedis() *redis.Client {
    if redisClient == nil {
        log.Fatal("Redis is not initialized. Call InitDatabase() first.")
    }
    return redisClient
}
```

4. Update your environment variables:
```bash
# Redis Configuration
export DB_DRIVER=redis
export REDIS_ADDR=localhost:6379
export REDIS_PASSWORD=
export REDIS_DB=0
```

5. Usage in your code:
```go
import "your-module/configs"

func YourFunction() {
    redis := configs.GetRedis()
    err := redis.Set(context.Background(), "key", "value", 0).Err()
    if err != nil {
        panic(err)
    }
}
```

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
- **Clean Structure**: Simple and straightforward organization
- **Easy to Extend**: Add your own models and features as needed
- **Database Flexibility**: Easy to switch between different database types
- **Future-Proof**: Ready to integrate with any database system
- **Thread-Safe**: Singleton pattern with sync.Once for safe initialization

## Adding New Features

### 1. Create Feature Structure
```
app/
├── http/
│   ├── controllers/
│   │   ├── interfaces/
│   │   │   └── userControllerInterface.go
│   │   └── userControllerImpl.go
├── repositories/
│   ├── interfaces/
│   │   └── userRepositoryInterface.go
│   └── userRepositoryImpl.go
└── services/
    ├── interfaces/
    │   └── userServiceInterface.go
    └── userServiceImpl.go
```

### 2. Register in Wire
In `app/wire/wire.go`:
```go
// Add new provider set
var userSet = wire.NewSet(
    repositories.NewUserRepository,
    services.NewUserService,
    controllers.NewUserController,
)

// Update InitializeControllers
func InitializeControllers() (
    interfaces.HealthController,
    interfaces.UserController,  // Add new controller
    error,
) {
    wire.Build(
        healthSet,
        userSet,  // Add new provider set
    )
    return nil, nil, nil
}
```

### 3. Add Routes
In `routes/api.go`:
```go
func SetupApiRoutes(app *fiber.App) {
    api := app.Group("/api")

    // Initialize controllers
    healthController, userController, err := wire.InitializeControllers()
    if err != nil {
        panic(err)
    }

    // Health routes
    api.Get("/health", healthController.Check)

    // User routes
    userGroup := api.Group("/users")
    userGroup.Post("/", userController.Create)
    userGroup.Get("/:id", userController.GetByID)
    userGroup.Get("/", userController.GetAll)
    userGroup.Put("/:id", userController.Update)
    userGroup.Delete("/:id", userController.Delete)
}
```

### 4. Generate Wire Code
After adding new feature, run:
```bash
wire ./app/wire
```

## Dependency Injection

This project uses Google Wire for dependency injection. The benefits include:

1. **Automatic Dependency Resolution**: Wire automatically resolves and injects dependencies
2. **Compile-time Safety**: All dependencies are checked at compile time
3. **Clean Architecture**: Clear separation of concerns
4. **Easy Testing**: Dependencies can be easily mocked
5. **Scalable**: Easy to add new features without changing existing code

### How It Works

1. **Provider Sets**: Each feature has its own provider set
2. **Wire Build**: Dependencies are resolved at compile time
3. **Interface-based**: All components are interface-based for better testing
4. **Modular**: Features are isolated and can be added/removed easily