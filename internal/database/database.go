package database

import (
    "database/sql"
    "log"

    _ "github.com/go-sql-driver/mysql"
)

// InitDB initializes the database connection
func InitDB() (*sql.DB, error) {
    // Connect to the database
    db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/your_database")
    if err != nil {
        log.Fatal("Error connecting to the database:", err)
        return nil, err
    }

    // Check database connection
    if err = db.Ping(); err != nil {
        log.Fatal("Error pinging database:", err)
        return nil, err
    }

    return db, nil
}

// Define database models here
