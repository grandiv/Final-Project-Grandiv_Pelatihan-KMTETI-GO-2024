package main

import (
	handler "final_project/api/handler"
	"final_project/src/db"
	"fmt"
	"log"
	"net/http"
)

func main() {
    // Initialize database connection
    database, err := db.DBConnection()
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // Create handler instances
    karyawanHandler := handler.NewKaryawanHandler(database)
    bukuHandler := handler.NewBukuHandler(database)

    // Setup router
    h := http.NewServeMux()

    // Register routes
    h.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, world!"))
    })
    h.HandleFunc("/api/buku/", bukuHandler.HandleBuku)
    h.HandleFunc("/api/karyawan/", karyawanHandler.HandleKaryawan)

    // Create server
    s := &http.Server{
        Addr:    ":8080",
        Handler: h,
    }

    // Start server
    fmt.Println("HTTP Server running on port 8080")
    if err := s.ListenAndServe(); err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}