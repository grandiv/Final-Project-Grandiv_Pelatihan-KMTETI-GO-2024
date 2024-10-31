package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"final_project/src/db"
	"final_project/src/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Service struct to encapsulate database operations
type KaryawanService struct {
    collection *mongo.Collection
}

// Constructor for KaryawanService
func NewKaryawanService(db *db.DB) *KaryawanService {
    return &KaryawanService{
        collection: db.MongoDB.Collection("karyawan"),
    }
}

type Karyawan struct {
    Nama         string            `json:"nama"`
    TanggalMasuk time.Time        `json:"tanggal_masuk"`
    StatusKerja  model.StatusKerja `json:"status_kerja"`
}

type KaryawanResponse struct {
    Data             []*Karyawan    `json:"data"`
}

type KaryawanRequest struct {
    Nama               string         `json:"nama"`
    NIK                string         `json:"nik"`
    PendidikanTerakhir string         `json:"pendidikan_terakhir"`
    TanggalMasuk       string         `json:"tanggal_masuk"` // Date format: YYYY-MM-DD
    StatusKerja        model.StatusKerja `json:"status_kerja"`
}

// Validation method for KaryawanRequest
func (req *KaryawanRequest) Validate() error {
    if req.Nama == "" {
        return fmt.Errorf("nama is required")
    }
    if req.NIK == "" {
        return fmt.Errorf("NIK is required")
    }
    if req.StatusKerja != model.StatusKontrak && req.StatusKerja != model.StatusTetap {
        return fmt.Errorf("invalid status kerja")
    }
    
    // Validate date format
    _, err := time.Parse("2006-01-02", req.TanggalMasuk)
    if err != nil {
        return fmt.Errorf("invalid date format, use YYYY-MM-DD")
    }
    
    return nil
}

// API 6: GetAllKaryawan (nama, tanggal masuk, status kerja)
func (s *KaryawanService) GetAllKaryawan(ctx context.Context) (*KaryawanResponse, error) {
    // Use context with timeout
    ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    cursor, err := s.collection.Find(ctx, bson.D{})
    if err != nil {
        return nil, fmt.Errorf("failed to retrieve karyawan: %v", err)
    }
    defer cursor.Close(ctx)

    var karyawanList []*Karyawan
	var modelKaryawanList []model.Karyawan
    if err = cursor.All(ctx, &modelKaryawanList); err != nil {
        return nil, fmt.Errorf("failed to decode karyawan list: %v", err)
    }

    for _, mk := range modelKaryawanList {
        karyawanList = append(karyawanList, &Karyawan{
            Nama:         mk.Nama,
            TanggalMasuk: mk.Tanggal_Masuk,
            StatusKerja:  mk.Status_Kerja,
        })
    }

    return &KaryawanResponse{
        Data: karyawanList,
    }, nil
}

// API 7: CreateKaryawan
func (s *KaryawanService) CreateKaryawan(ctx context.Context, req io.Reader) error {
    ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    var karyawanReq KaryawanRequest
    if err := json.NewDecoder(req).Decode(&karyawanReq); err != nil {
        return fmt.Errorf("invalid request body: %v", err)
    }

    if err := karyawanReq.Validate(); err != nil {
        return err
    }

    // Parse date string to time.Time
    tanggalMasuk, _ := time.Parse("2006-01-02", karyawanReq.TanggalMasuk)

    karyawan := model.Karyawan{
        Id:                  primitive.NewObjectID(),
        Nama:                karyawanReq.Nama,
        NIK:                 karyawanReq.NIK,
        Pendidikan_Terakhir: karyawanReq.PendidikanTerakhir,
        Tanggal_Masuk:       tanggalMasuk,
        Status_Kerja:        karyawanReq.StatusKerja,
    }

    _, err := s.collection.InsertOne(ctx, karyawan)
    if err != nil {
        return fmt.Errorf("failed to create karyawan: %v", err)
    }

    return nil
}