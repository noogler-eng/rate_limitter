# Rate Limiter with Redis

This repository implements a simple rate limiter using Redis as the storage backend. It provides a middleware that can be used to protect your HTTP endpoints from excessive requests.

## Features

- **Redis-backed:** Uses Redis for fast and efficient storage of request timestamps.
- **Configurable Limits:** Easily configure the request limit and time window.
- **IP-based Rate Limiting:** Limits requests based on the client's IP address.
- **Middleware Implementation:** Provides a middleware function that can be easily integrated with your HTTP handlers.

## Prerequisites

- Go 1.16 or later
- Redis server

## Installation

1.  **Clone the repository:**

    ```bash
    git clone <repository_url>
    cd rate-limiter
    ```

2.  **Install dependencies:**

    ```bash
    go mod tidy
    ```

3.  **Ensure Redis is running:**

    Make sure you have a Redis server running and accessible.

## Usage

1.  **Initialize Redis:**

    In your `main.go` file, initialize the Redis client using the `redisdb.InitRedis()` function:

    ```go
    package main

    import (
        "fmt"
        "net/http"

        "[github.com/noogler-eng/rate-limiter/limitter](https://www.google.com/search?q=https://github.com/noogler-eng/rate-limiter/limitter)"
        "[github.com/noogler-eng/rate-limiter/redisdb](https://www.google.com/search?q=https://github.com/noogler-eng/rate-limiter/redisdb)"
    )

    func main() {
        redisdb.InitRedis()
        defer redisdb.RedisClient.Close()

        // ... your HTTP handlers ...
    }
    ```

2.  **Apply the rate limiter middleware:**

    Wrap your HTTP handler with the `limitter.RateLimitter()` function:

    ```go
    package main

    import (
        "encoding/json"
        "fmt"
        "net/http"

        "[github.com/noogler-eng/rate-limiter/limitter](https://www.google.com/search?q=https://github.com/noogler-eng/rate-limiter/limitter)"
        "[github.com/noogler-eng/rate-limiter/redisdb](https://www.google.com/search?q=https://github.com/noogler-eng/rate-limiter/redisdb)"
    )

    func main() {
        redisdb.InitRedis()
        defer redisdb.RedisClient.Close()

        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
            var res_string string = "this is route for testing"
            json.NewEncoder(w).Encode(res_string)
        })

        myHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            w.Write([]byte("This is a rate-limited endpoint"))
        })

        http.HandleFunc("/ping", limitter.RateLimitter(myHandler))

        fmt.Println("Server running on port 8080...")
        http.ListenAndServe(":8080", nil)
    }
    ```

3.  **Configure limits:**

    Modify the `LimitInMinutes` and `RequestLimit` variables in the `limitter/limitter.go` file to adjust the rate limiting behavior:

    ```go
    package limitter

    // ...

    var LimitInMinutes int64 = 1 // Limit window in minutes
    var RequestLimit int64 = 5   // Maximum number of requests within the window

    // ...
    ```

4.  **Run the application:**

    ```bash
    go run main.go
    ```

## Redis Configuration

The Redis client is initialized with default options. If you need to customize the Redis connection, you can modify the `redis.NewClient()` call in `redisdb/redisdb.go`.

```go
package redisdb

import (
        "context"
        "fmt"
        "log"

        "[github.com/redis/go-redis/v9](https://www.google.com/search?q=https://github.com/redis/go-redis/v9)"
)

var RedisClient *redis.Client

func InitRedis() {
        RedisClient = redis.NewClient(&redis.Options{
                Addr:     "localhost:6379", // Customize Redis address
                Password: "",               // Customize Redis password
                DB:       0,                // Customize Redis database
        })
        fmt.Println("redis has been initlized!")

        _, err := RedisClient.Ping(context.Background()).Result()
        if err != nil {
                log.Fatalf("Could not connect to Redis: %v", err)
        }
}
```

## Testing

You can use `curl` or any other HTTP client to test the rate limiter:

```bash
curl http://localhost:8080/ping
```

## Error Handling

The rate limiter handles Redis errors by returning a 500 Internal Server Error response. When the rate limit is exceeded, it returns a 429 Too Many Requests response with a JSON error message.

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue.