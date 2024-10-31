package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StatusKerja string

const (
    StatusKontrak StatusKerja = "Kontrak"
    StatusTetap StatusKerja = "Tetap"
)

type Karyawan struct {
    Id                  primitive.ObjectID `bson:"_id,omitempty"`
    Nama                string            `bson:"nama"`
    NIK                 string            `bson:"nik"`
    Pendidikan_Terakhir string            `bson:"pendidikan_terakhir"`
    Tanggal_Masuk       time.Time         `bson:"tanggal_masuk"`
    Status_Kerja        StatusKerja       `bson:"status_kerja"`
}