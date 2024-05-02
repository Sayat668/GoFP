package main

import (
    "fmt"
    "log"
    "net/http"

    "GOFP/internal/auth"
    "GOFP/internal/order"
    "GOFP/internal/menu"
    "GOFP/internal/analytics"
    "GOFP/internal/database"
)

func main() {
    // Initialize database connection
    db, err := database.InitDB()
    if err != nil {
        log.Fatal("Error initializing database:", err)
    }
    defer db.Close()

    // Initialize HTTP routes
    http.HandleFunc("/login", auth.LoginHandler)
    http.HandleFunc("/orders", order.OrdersHandler)
    http.HandleFunc("/menu", menu.MenuHandler)
    http.HandleFunc("/analytics", analytics.AnalyticsHandler)

    // Start the HTTP server
    fmt.Println("Server running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
