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

type Karyawan struct {
	Nama string `json:"nama"`
	NIK string `json:"nik"`
	Pendidikan_Terakhir string `json:"pendidikan_terakhir"`
	Tanggal_Masuk string `json:"tanggal_masuk"`
	Status_Kerja model.StatusKerja `json:"status_kerja"`
}

type KaryawanResponse struct {
	Data []*Karyawan `json:"data"`
}

// KaryawanRequest is a struct to hold request data
type KaryawanRequest struct {
	Nama string `json:"nama"`
	NIK string `json:"nik"`
	Pendidikan_Terakhir string `json:"pendidikan_terakhir"`
	Tanggal_Masuk string `json:"tanggal_masuk"`
	Status_Kerja model.StatusKerja `json:"status_kerja"`
}

// API 6: Get All Karyawan (nama, tanggal masuk, status kerja)
func GetAllKaryawan() (*KaryawanResponse, error) {
	db, err := db.DBConnection()
	if err != nil {
		return nil, errors.New("internal server error")
	} 
	defer db.MongoDB.Client().Disconnect(context.TODO())

	coll := db.MongoDB.Collection("karyawan")
	cur, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, errors.New("internal server error")
	}

	var karyawanList []*Karyawan

	for cur.Next(context.TODO()) {
		var karyawan model.Karyawan
		cur.Decode(&karyawan)
		karyawanList = append(karyawanList, &Karyawan{
			Nama: karyawan.Nama,
			Tanggal_Masuk: karyawan.Tanggal_Masuk,
			Status_Kerja: karyawan.Status_Kerja,
		})
	}
	return &KaryawanResponse{
		Data: karyawanList,
	}, nil
}

// API 7: Create Karyawan
func CreateKaryawan(req io.Reader) error {
	var karyawanReq KaryawanRequest
	err := json.NewDecoder(req).Decode(&karyawanReq)
	if err != nil {
		return errors.New("bad request")
	}

	db, err := db.DBConnection()
	if err != nil {
		return errors.New("internal server error")
	}
	defer db.MongoDB.Client().Disconnect(context.TODO())

	coll := db.MongoDB.Collection("karyawan")
	_, err = coll.InsertOne(context.TODO(), model.Karyawan{
		Id: primitive.NewObjectID(),
		Nama: karyawanReq.Nama,
		NIK: karyawanReq.NIK,
		Pendidikan_Terakhir: karyawanReq.Pendidikan_Terakhir,
		Tanggal_Masuk: karyawanReq.Tanggal_Masuk,
		Status_Kerja: karyawanReq.Status_Kerja,
	})
	if err != nil {
		return errors.New("internal server error")
	}
	return nil
}