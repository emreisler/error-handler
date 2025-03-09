# Go Error Handler Middleware for Gin

A lightweight and efficient **error handling middleware** for **Gin-based web services** in Go. This library provides a **centralized error management system** that:
- Catches **custom service errors** and maps them to appropriate **HTTP status codes**.
- Automatically detects **database errors** (GORM, sqlx, etc.).
- Recovers from **panics**, preventing server crashes.
- Simplifies error handling so you **only need to call `c.Error(err)`** in handlers.

---

## **🚀 Features**
✅ **Centralized error handling** – No need for manual `errors.As()` checks.  
✅ **Automatic database error detection** – Handles `gorm.ErrRecordNotFound`, `sql.ErrNoRows`, and more.  
✅ **Panic recovery** – Prevents server crashes from unexpected panics.  
✅ **Consistent HTTP responses** – Every error maps to a structured JSON response.  
✅ **Minimal integration effort** – Plug & play with Gin middleware.

---

## **📦 Installation**
```sh
go get github.com/emreisler/gin-error-handler
```

---

## **📖 Usage Guide**

### **Setup in a Gin Application**
```go
package main

import (
    "github.com/emreisler/gin-error-handler"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.New()
    r.Use(ginerrorhandler.ErrorHandler())
    
    // Define your routes here
    
    r.Run()
}
```

### **How to Use in Request Handlers**
In your request handlers, you can use the middleware like this:
```go
func MyHandler(c *gin.Context) {
    err := SomeFunction()
    if err != nil {
        c.Error(err)  // This will be handled by the middleware
        return
    }
    c.JSON(200, gin.H{"message": "success"})
}
```

### **Automatic Database Error Handling**
The middleware automatically detects and handles database errors from GORM and sqlx, such as:
- `gorm.ErrRecordNotFound`
- `sql.ErrNoRows`

### **Panic Recovery Example**
If your application encounters a panic, the middleware will recover and return a 500 status code:
```go
func PanicHandler(c *gin.Context) {
    defer func() {
        if r := recover(); r != nil {
            c.Error(fmt.Errorf("panic occurred: %v", r))
        }
    }()
    // Your code that may panic
}
```

---

## **🛠 Error Types**
| Error Type                     | HTTP Status Code |
|--------------------------------|------------------|
| Custom Service Error           | 400              |
| Not Found                      | 404              |
| Internal Server Error          | 500              |

---

## **💾 Database Error Handling**
The middleware provides automatic detection of common database errors, allowing you to focus on your application logic without worrying about error handling.

---

## **🚨 Panic Recovery**
The middleware is designed to catch panics and prevent server crashes, ensuring that your application remains stable even in the face of unexpected errors.

---

## **🤝 Why Use This Library**
Integrating this library simplifies your error handling process, reduces boilerplate code, and enhances the stability of your Gin-based applications.

---

## **💡 Contribution Guidelines**
We welcome contributions! Please fork the repository and submit a pull request with your changes.

---

## **📄 License**
This project is licensed under the MIT License. See the LICENSE file for details.