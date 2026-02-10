package main

import (
    "fmt"
    "net/http"
)

func main() {
    // 1. Define a simple "home" route
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Waste Logistics Engine API is running!")
    })

    // 2. Start the server on port 8080
    fmt.Println("🚀 Server starting on http://localhost:8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Error starting server:", err)
    }
}