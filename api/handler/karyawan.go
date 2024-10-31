package handler

import (
	"encoding/json"
	"net/http"

	"final_project/src/db"
	"final_project/src/service"
)

// KaryawanHandler handles employee-related HTTP requests
type KaryawanHandler struct {
    service *service.KaryawanService
}

// NewKaryawanHandler creates a new handler with the given service
func NewKaryawanHandler(db *db.DB) *KaryawanHandler {
    return &KaryawanHandler{
        service: service.NewKaryawanService(db),
    }
}

// HandleKaryawan handles GET and POST requests
func (h *KaryawanHandler) HandleKaryawan(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        h.getAllKaryawan(w, r)
    case http.MethodPost:
        h.createKaryawan(w, r)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

// getAllKaryawan retrieves all employees
func (h *KaryawanHandler) getAllKaryawan(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    
    karyawanList, err := h.service.GetAllKaryawan(ctx)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(karyawanList)
}

// createKaryawan creates a new employee
func (h *KaryawanHandler) createKaryawan(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()

    if err := h.service.CreateKaryawan(ctx, r.Body); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    w.WriteHeader(http.StatusCreated)
    w.Write([]byte("Karyawan created successfully"))
}