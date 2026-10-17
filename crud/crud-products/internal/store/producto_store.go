package store

import (
	"context"
	"crud-products/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductoStore struct {
	collection *mongo.Collection
}

func NewProductoStore(db *mongo.Database) *ProductoStore {
	return &ProductoStore{
		collection: db.Collection("productos"),
	}
}

func (ps *ProductoStore) Crear(prd model.Producto) error {
	_, err := ps.collection.InsertOne(context.TODO(), prd)
	if err != nil {
		return err
	}
	return nil
}

func (ps *ProductoStore) ObtenerTodos() ([]model.Producto, error) {
	productos := []model.Producto{}
	product, err := ps.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	err = product.All(context.TODO(), &productos)
	if err != nil {
		return nil, err
	}
	return productos, nil
}
func (ps *ProductoStore) ObtenerPorID(id string) (model.Producto, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Producto{}, err
	}
	resultado := ps.collection.FindOne(context.TODO(), bson.M{"_id": objectID})
	producto := model.Producto{}
	err = resultado.Decode(&producto)
	if err != nil {
		return model.Producto{}, err
	}
	return producto, nil
}
func (ps *ProductoStore) Actualizar(id string, producto model.Producto) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	objetoActualizar := bson.M{"_id": objectID}
	campoActualizar := bson.M{"$set": producto}
	_, err = ps.collection.UpdateOne(context.TODO(), objetoActualizar, campoActualizar)
	if err != nil {
		return err
	}
	return nil
}
func (ps *ProductoStore) Eliminar(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = ps.collection.DeleteOne(context.TODO(), bson.M{"_id": objectID})
	if err != nil {
		return err
	}
	return nil
}
