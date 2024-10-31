package handler

import (
	"encoding/json"
	"net/http"

	"final_project/src/service"
)

// BukuHandler is a handler to handle request to /api/buku
func BukuHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		data, err := service.GetAllBuku()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
		return

	case "POST":
		err := service.CreateBuku(r.Body)
		if err != nil {
			if err.Error() == "Bad Request" {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("Product Created Successfully")
		return

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}
}