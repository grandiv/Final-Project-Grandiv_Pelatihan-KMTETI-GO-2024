package api

import (
	"final_project/api/handler"
	"final_project/src/db"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    // Initialize database connection
    database, err := db.DBConnection()
    if err != nil {
        http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
        return
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

    // Serve the request
    h.ServeHTTP(w, r)
}