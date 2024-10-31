package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"final_project/src/db"
	"final_project/src/service"
)

// BukuHandler handles book-related HTTP requests
type BukuHandler struct {
    service *service.BukuService
}

// NewBukuHandler creates a new handler with the given service
func NewBukuHandler(db *db.DB) *BukuHandler {
    return &BukuHandler{
        service: service.NewBukuService(db),
    }
}

// HandleBuku handles all book-related requests
func (h *BukuHandler) HandleBuku(w http.ResponseWriter, r *http.Request) {
    // Check if we have an ID in the path
    path := strings.TrimPrefix(r.URL.Path, "/api/buku/")
    if path != r.URL.Path { // We have a path parameter
        // Only handle ID-based operations if there's actually an ID
        if path != "" {
            h.handleBukuWithID(w, r, path)
            return
        }
    }

    // Handle non-ID endpoints
    switch r.Method {
    case http.MethodGet:
        h.getAllBuku(w, r)
    case http.MethodPost:
        h.createBuku(w, r)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func (h *BukuHandler) handleBukuWithID(w http.ResponseWriter, r *http.Request, id string) {
    switch r.Method {
    case http.MethodGet:
        h.getBukuByID(w, r, id)
    case http.MethodPut:
        h.updateBuku(w, r, id)
    case http.MethodDelete:
        h.deleteBuku(w, r, id)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func (h *BukuHandler) getAllBuku(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    bukuList, err := h.service.GetAllBuku(ctx)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(bukuList)
}

func (h *BukuHandler) getBukuByID(w http.ResponseWriter, r *http.Request, id string) {
    ctx := r.Context()
    buku, err := h.service.GetBukuByID(ctx, id)
    if err != nil {
        status := http.StatusInternalServerError
        if err.Error() == "buku not found" {
            status = http.StatusNotFound
        }
        http.Error(w, err.Error(), status)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(buku)
}

func (h *BukuHandler) createBuku(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    if err := h.service.CreateBuku(ctx, r.Body); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    w.WriteHeader(http.StatusCreated)
    w.Write([]byte("Buku created successfully"))
}

func (h *BukuHandler) updateBuku(w http.ResponseWriter, r *http.Request, id string) {
    ctx := r.Context()
    if err := h.service.UpdateBuku(ctx, id, r.Body); err != nil {
        status := http.StatusInternalServerError
        if err.Error() == "buku not found" {
            status = http.StatusNotFound
        }
        http.Error(w, err.Error(), status)
        return
    }
    w.Write([]byte("Buku updated successfully"))
}

func (h *BukuHandler) deleteBuku(w http.ResponseWriter, r *http.Request, id string) {
    ctx := r.Context()
    if err := h.service.DeleteBuku(ctx, id); err != nil {
        status := http.StatusInternalServerError
        if err.Error() == "buku not found" {
            status = http.StatusNotFound
        }
        http.Error(w, err.Error(), status)
        return
    }
    w.Write([]byte("Buku deleted successfully"))
}