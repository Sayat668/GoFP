package menu

import (
    "net/http"
    "encoding/json"
    "GOFP/internal/database"
)

// MenuItem represents a menu item in the canteen
type MenuItem struct {
    ID    int     `json:"id"`
    Name  string  `json:"name"`
    Price float64 `json:"price"`
    // Add other menu item fields as needed
}

// CreateMenuItemHandler handles requests to add new menu items
func CreateMenuItemHandler(w http.ResponseWriter, r *http.Request) {
    // Example: Parse menu item data from request body
    var newItem MenuItem
    err := json.NewDecoder(r.Body).Decode(&newItem)
    if err != nil {
        http.Error(w, "Invalid menu item data", http.StatusBadRequest)
        return
    }

    // Example: Save the new menu item to the database
    savedItem, err := database.CreateMenuItem(newItem)
    if err != nil {
        http.Error(w, "Error creating menu item", http.StatusInternalServerError)
        return
    }

    // Respond with the created menu item
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(savedItem)
}
