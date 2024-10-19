package handler

import (
	"encoding/json"
	"net/http"
)

type Buku struct {
	Id    int    `json:"id"`
	Judul string `json:"judul"`
	Penulis string `json:"penulis"`
	Tahun int    `json:"tahun"`
	Stok uint8  `json:"stok"`
	Harga int    `json:"harga"`
}

var BukuList []*Buku = []*Buku{
	{
		Id:   1,
		Judul: "Buku A",
		Penulis: "Penulis A",
		Tahun: 2021,
		Stok: 10,
		Harga: 10000,
	},
	{
		Id:   2,
		Judul: "Buku B",
		Penulis: "Penulis B",
		Tahun: 2022,
		Stok: 20,
		Harga: 20000,
	},
}

// BukuRequest is a struct to hold request data
type BukuRequest struct {
	Judul string `json:"judul"`
	Penulis string `json:"penulis"`
	Tahun int    `json:"tahun"`
	Stok uint8  `json:"stok"`
	Harga int    `json:"harga"`
}

// BukuHandler is a handler to handle request to /api/buku
func BukuHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(BukuList)
		return

	case "POST":
		var data BukuRequest

		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		b := Buku{
			Id:   len(BukuList) + 1,
			Judul: data.Judul,
			Penulis: data.Penulis,
			Tahun: data.Tahun,
			Stok: data.Stok,
			Harga: data.Harga,
		}

		BukuList = append(BukuList, &b)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Book Added Successfully"))

		return

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}
}