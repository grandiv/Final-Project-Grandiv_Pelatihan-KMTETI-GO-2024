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
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Service struct to encapsulate database operations
type BukuService struct {
    collection *mongo.Collection
}

// Constructor for BukuService
func NewBukuService(db *db.DB) *BukuService {
    return &BukuService{
        collection: db.MongoDB.Collection("buku"),
    }
}

type Buku struct {
    Judul   string `json:"judul"`
    Penulis string `json:"penulis"`
    Tahun   int    `json:"tahun"`
    Stok    uint8  `json:"stok"`
    Harga   int    `json:"harga"`
}

type BukuResponse struct {
    Data []*Buku `json:"data"`
}

type BukuGetAll struct {
	Judul   string `json:"judul"`
	Penulis string `json:"penulis"`
	Harga   int    `json:"harga"`
}

type BukuResponseGetAll struct {
	Data []*BukuGetAll `json:"data"`
}

type BukuRequest struct {
    Judul   string `json:"judul"`
    Penulis string `json:"penulis"`
    Tahun   int    `json:"tahun"`
    Stok    uint8  `json:"stok"`
    Harga   int    `json:"harga"`
}

type UpdateBukuRequest struct {
    Stok  uint8 `json:"stok"`
    Harga int   `json:"harga"`
}

// Validate validates the buku request
func (req *BukuRequest) Validate() error {
    if req.Judul == "" {
        return fmt.Errorf("judul is required")
    }
    if req.Penulis == "" {
        return fmt.Errorf("penulis is required")
    }
    if req.Tahun <= 0 {
        return fmt.Errorf("tahun must be positive")
    }
    if req.Harga <= 0 {
        return fmt.Errorf("harga must be positive")
    }
    return nil
}

// API 1: GetAllBuku (judul, penulis, harga)
func (s *BukuService) GetAllBuku(ctx context.Context) (*BukuResponseGetAll, error) {
    ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    // Define projection to only get required fields
    projection := bson.D{
        {Key: "judul", Value: 1},
        {Key: "penulis", Value: 1},
        {Key: "harga", Value: 1},
    }

    cursor, err := s.collection.Find(ctx, bson.D{}, &options.FindOptions{
        Projection: projection,
    })
    if err != nil {
        return nil, fmt.Errorf("failed to retrieve buku: %v", err)
    }
    defer cursor.Close(ctx)

    var modelBukuList []model.Buku
    if err = cursor.All(ctx, &modelBukuList); err != nil {
        return nil, fmt.Errorf("failed to decode buku list: %v", err)
    }

    var bukuList []*BukuGetAll
    for _, mk := range modelBukuList {
        bukuList = append(bukuList, &BukuGetAll{
            Judul:   mk.Judul,
            Penulis: mk.Penulis,
            Harga:   mk.Harga,
        })
    }

    return &BukuResponseGetAll{Data: bukuList}, nil
}

// API 2: GetBukuByID (judul, penulis, tahun terbit, stok, harga)
func (s *BukuService) GetBukuByID(ctx context.Context, id string) (*Buku, error) {
    ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, fmt.Errorf("invalid ID format: %v", err)
    }

    // Define projection to only get required fields
    projection := bson.D{
        {Key: "judul", Value: 1},
        {Key: "penulis", Value: 1},
        {Key: "tahun", Value: 1},
        {Key: "harga", Value: 1},
        {Key: "stok", Value: 1},
    }

    var buku model.Buku
    err = s.collection.FindOne(ctx, bson.M{"_id": objectID}, &options.FindOneOptions{
        Projection: projection,
    }).Decode(&buku)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return nil, fmt.Errorf("buku not found")
        }
        return nil, fmt.Errorf("failed to retrieve buku: %v", err)
    }

    return &Buku{
        Judul:   buku.Judul,
        Penulis: buku.Penulis,
        Tahun:   buku.Tahun,
        Stok:    buku.Stok,
        Harga:   buku.Harga,
    }, nil
}

// API 3: CreateBuku 
func (s *BukuService) CreateBuku(ctx context.Context, req io.Reader) error {
    ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    var bukuReq BukuRequest
    if err := json.NewDecoder(req).Decode(&bukuReq); err != nil {
        return fmt.Errorf("invalid request body: %v", err)
    }

    if err := bukuReq.Validate(); err != nil {
        return err
    }

    buku := model.Buku{
        Id:      primitive.NewObjectID(),
        Judul:   bukuReq.Judul,
        Penulis: bukuReq.Penulis,
        Tahun:   bukuReq.Tahun,
        Stok:    bukuReq.Stok,
        Harga:   bukuReq.Harga,
    }

    _, err := s.collection.InsertOne(ctx, buku)
    if err != nil {
        return fmt.Errorf("failed to create buku: %v", err)
    }

    return nil
}

// API 4: UpdateBuku (stok dan harga)
func (s *BukuService) UpdateBuku(ctx context.Context, id string, req io.Reader) error {
    ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    var updateReq UpdateBukuRequest
    if err := json.NewDecoder(req).Decode(&updateReq); err != nil {
        return fmt.Errorf("invalid request body: %v", err)
    }

    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return fmt.Errorf("invalid ID format: %v", err)
    }

    update := bson.M{
        "$set": bson.M{
            "stok":  updateReq.Stok,
            "harga": updateReq.Harga,
        },
    }

    result, err := s.collection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
    if err != nil {
        return fmt.Errorf("failed to update buku: %v", err)
    }

    if result.MatchedCount == 0 {
        return fmt.Errorf("buku not found")
    }

    return nil
}

// API 5: DeleteBuku 
func (s *BukuService) DeleteBuku(ctx context.Context, id string) error {
    ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return fmt.Errorf("invalid ID format: %v", err)
    }

    result, err := s.collection.DeleteOne(ctx, bson.M{"_id": objectID})
    if err != nil {
        return fmt.Errorf("failed to delete buku: %v", err)
    }

    if result.DeletedCount == 0 {
        return fmt.Errorf("buku not found")
    }

    return nil
}