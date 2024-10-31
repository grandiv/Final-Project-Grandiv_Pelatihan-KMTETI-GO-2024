package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Buku struct {
	Id    primitive.ObjectID    `bson:"_id,omitempty"`
	Judul string `bson:"judul"`
	Penulis string `bson:"penulis"`
	Tahun int    `bson:"tahun"`
	Stok uint8  `bson:"stok"`
	Harga int    `bson:"harga"`
}