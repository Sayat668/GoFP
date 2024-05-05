package menu

import (
    "net/http"
)

// MenuHandler handles requests related to the canteen menu
func MenuHandler(w http.ResponseWriter, r *http.Request) {
    // Example: Handle different menu-related actions based on HTTP method
    switch r.Method {
    case http.MethodPost:
        CreateMenuItemHandler(w, r)
    // Add more cases for other menu actions like update, delete, etc.
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}
