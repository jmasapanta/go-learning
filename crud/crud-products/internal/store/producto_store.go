package store

import (
	"crud/crud-products/internal/model"
	"mgo.mongodb.org/mongo-driver/bson"
)

type ProductoStore struct{
	collection *mongo.Collection
}
func NewProductoStore(db *mongo.Database) *ProductoStore{
	return &ProductoStore{
		collection: db.Collection("productos")
	}
}