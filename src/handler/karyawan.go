package handler

import (
	"encoding/json"
	"net/http"
)

type StatusKerja string

const (
	StatusKontrak StatusKerja = "Kontrak"
	StatusTetap StatusKerja = "Tetap"
)

type Karyawan struct {
	Id    int    `json:"id"`
	Nama string `json:"nama"`
	NIK string `json:"nik"`
	Pendidikan_Terakhir string `json:"pendidikan_terakhir"`
	Tanggal_Masuk string `json:"tanggal_masuk"`
	Status_Kerja StatusKerja `json:"status_kerja"`

}

var KaryawanList []*Karyawan = []*Karyawan{
	{
		Id:   1,
		Nama: "Karyawan A",
		NIK: "123456",
		Pendidikan_Terakhir: "S1",
		Tanggal_Masuk: "2021-01-01",
		Status_Kerja: StatusKontrak,
	},
	{
		Id:   2,
		Nama: "Karyawan B",
		NIK: "123457",
		Pendidikan_Terakhir: "S2",
		Tanggal_Masuk: "2021-01-02",
		Status_Kerja: StatusTetap,
	},
}

// KaryawanRequest is a struct to hold request data
type KaryawanRequest struct {
	Nama string `json:"nama"`
	NIK string `json:"nik"`
	Pendidikan_Terakhir string `json:"pendidikan_terakhir"`
	Tanggal_Masuk string `json:"tanggal_masuk"`
	Status_Kerja StatusKerja `json:"status_kerja"`
}

// KaryawanHandler is a handler to handle request to /api/buku
func KaryawanHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(BukuList)
		return

	case "POST":
		var data KaryawanRequest

		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		if data.Status_Kerja != StatusKontrak && data.Status_Kerja != StatusTetap {
        http.Error(w, "Invalid Status_Kerja. Must be 'Kontrak' or 'Tetap'.", http.StatusBadRequest)
        return
		}

		k := Karyawan{
			Id:   len(BukuList) + 1,
			Nama: data.Nama,
			NIK: data.NIK,
			Pendidikan_Terakhir: data.Pendidikan_Terakhir,
			Tanggal_Masuk: data.Tanggal_Masuk,
			Status_Kerja: data.Status_Kerja,
		}

		KaryawanList = append(KaryawanList, &k)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Karyawan Added Successfully"))

		return

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}
}