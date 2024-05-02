package database

// Define your database models here
type User struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
    Password string `json:"password"`
    // Add other user-related fields as needed
}

type Order struct {
    ID       int     `json:"id"`
    UserID   int     `json:"user_id"`
    MenuID   int     `json:"menu_id"`
    Quantity int     `json:"quantity"`
    // Add other order-related fields as needed
}

type MenuItem struct {
    ID    int     `json:"id"`
    Name  string  `json:"name"`
    Price float64 `json:"price"`
    // Add other menu item fields as needed
}
