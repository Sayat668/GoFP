package order

import (
    "net/http"
    "encoding/json"
    "GOFP/internal/database"
)

// Order represents the order model in the database
type Order struct {
    ID       int `json:"id"`
    UserID   int `json:"user_id"`
    MenuID   int `json:"menu_id"`
    Quantity int `json:"quantity"`
    // Add other order-related fields as needed
}

// CreateOrderHandler handles requests to create new orders
func CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
    // Example: Parse order data from request body
    var newOrder Order
    err := json.NewDecoder(r.Body).Decode(&newOrder)
    if err != nil {
        http.Error(w, "Invalid order data", http.StatusBadRequest)
        return
    }

    // Example: Save the new order to the database
    savedOrder, err := database.CreateOrder(newOrder)
    if err != nil {
        http.Error(w, "Error creating order", http.StatusInternalServerError)
        return
    }

    // Respond with the created order
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(savedOrder)
}
