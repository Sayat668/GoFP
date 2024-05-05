package auth

import (
    "net/http"
    "GOFP/internal/database"
)

// UserLogin handles user login requests
func UserLogin(w http.ResponseWriter, r *http.Request) {
    // Example: Validate user credentials and generate JWT token
    // Retrieve username and password from request body
    username := r.FormValue("username")
    password := r.FormValue("password")

    // Example: Validate username and password against database
    user, err := database.GetUserByUsername(username)
    if err != nil {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }

    // Example: Compare hashed password
    if user.Password != hashPassword(password) {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }

    // Generate JWT token
    token, err := generateToken(user.ID)
    if err != nil {
        http.Error(w, "Error generating token", http.StatusInternalServerError)
        return
    }

    // Respond with the generated token
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"token": token})
}

// Middleware function to authenticate requests using JWT token
func Authenticate(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Example: Extract JWT token from request header and validate
        token := extractTokenFromHeader(r)
        if token == "" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        // Example: Verify JWT token
        userID, err := verifyToken(token)
        if err != nil {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        // Set user ID in context for further processing
        ctx := context.WithValue(r.Context(), "userID", userID)
        r = r.WithContext(ctx)

        // Call the next handler
        next.ServeHTTP(w, r)
    })
}
