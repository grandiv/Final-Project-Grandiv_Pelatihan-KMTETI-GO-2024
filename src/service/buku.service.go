package service

import (
	"context"
	"encoding/json"
	"errors"
	"io"

	"final_project/src/db"
	"final_project/src/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Buku struct {
	Judul string `json:"judul"`
	Penulis string `json:"penulis"`
	Tahun int `json:"tahun"`
	Stok uint8 `json:"stok"`
	Harga int `json:"harga"`
}

type BukuResponse struct {
	Data []*Buku `json:"data"`
}

// BukuRequest is a struct to hold request data
type BukuRequest struct {
	Judul string `json:"judul"`
	Penulis string `json:"penulis"`
	Tahun int    `json:"tahun"`
	Stok uint8  `json:"stok"`
	Harga int    `json:"harga"`
}

// API 1: Get All Buku (judul, penulis, harga)
func GetAllBuku() (*BukuResponse, error) {
	db, err := db.DBConnection()
	if err != nil {
		return nil, errors.New("internal server error")
	}
	defer db.MongoDB.Client().Disconnect(context.TODO())

	coll := db.MongoDB.Collection("buku")
	cur, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, errors.New("internal server error")
	}

	var bukuList []*Buku 

	for cur.Next(context.TODO()) {
		var buku model.Buku 
		cur.Decode(&buku)
		bukuList = append(bukuList, &Buku{
			Judul: buku.Judul,
			Penulis: buku.Penulis,
			Harga: buku.Harga,
		})
	}
	return &BukuResponse{
		Data: bukuList,
	}, nil
}

// API 2: Get Buku by ID (judul, penulis, tahun, stok, harga)

// API 3: Create Buku
func CreateBuku(req io.Reader) error {
	var bukuReq BukuRequest 
	err := json.NewDecoder(req).Decode(&bukuReq)
	if err != nil {
		return errors.New("bad request")
	}

	db, err := db.DBConnection()
	if err != nil {
		return errors.New("internal server error")
	}
	defer db.MongoDB.Client().Disconnect(context.TODO())

	coll := db.MongoDB.Collection("buku")
	_, err = coll.InsertOne(context.TODO(), model.Buku {
		Id: primitive.NewObjectID(),
		Judul: bukuReq.Judul,
		Penulis: bukuReq.Penulis,
		Tahun: bukuReq.Tahun,
		Stok: bukuReq.Stok,
		Harga: bukuReq.Harga,
	})
	if err != nil {
		return errors.New("internal server error")
	}
	return nil
}

// API 4: Update Buku (stok dan harga)

// API 5: Delete Buku