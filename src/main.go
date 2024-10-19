package main

import (
	"final_project/src/handler"
	"fmt"
	"net/http"
)

func main() {
	h := http.NewServeMux()

	s := &http.Server{
		Addr: ":8080",
		Handler: h,
	}

	h.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})
	h.HandleFunc("/api/buku", handler.BukuHandler)
	h.HandleFunc("/api/karyawan", handler.KaryawanHandler)

	fmt.Println("HTTP Server running on port 8080")
	err := s.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
}