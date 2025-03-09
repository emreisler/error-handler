# Error Handler Library

## Introduction
The Error Handler Library provides a centralized solution for managing errors in Go applications. Its primary purpose is to streamline error handling across various frameworks, ensuring that developers can focus on building features rather than managing error states.

## 🚀 Features
- Centralized error handling
- Automatic database error detection
- Panic recovery

## 📦 Installation
```sh
go get github.com/emreisler/error-handler
```

## 🛠 Usage Examples
### Gin Middleware
```go
import "github.com/emreisler/error-handler"

func main() {
    r := gin.Default()
    r.Use(error_handler.GinMiddleware())
    // Your routes here
}
```

### net/http Middleware
```go
import "github.com/emreisler/error-handler"

func main() {
    http.Handle("/", error_handler.NetHTTPMiddleware(http.HandlerFunc(yourHandler)))
    http.ListenAndServe(":8080", nil)
}
```

### Chi Middleware
```go
import "github.com/emreisler/error-handler"

func main() {
    r := chi.NewRouter()
    r.Use(error_handler.ChiMiddleware())
    // Your routes here
}
```

### Fiber Middleware
```go
import "github.com/emreisler/error-handler"

func main() {
    app := fiber.New()
    app.Use(error_handler.FiberMiddleware())
    // Your routes here
}
```

## 🔗 Database Error Handling
The `mapDBError` function automatically maps database errors to a standardized error type, allowing for consistent error responses.

## ⚠️ Panic Recovery
The library includes built-in panic recovery to prevent application crashes due to unexpected errors.

## 📌 Error Type Mapping
The library provides a mechanism for mapping different error types to a unified structure, making it easier to handle errors across various components.

## 🤝 Contribution Guidelines
We welcome contributions! Please submit a pull request or open an issue for any enhancements or bug fixes.

## 📜 License Information
This library is licensed under the MIT License. See the LICENSE file for more details.
