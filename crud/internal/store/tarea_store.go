package store

import (
	"context"
	"crud/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TareaStore struct {
	collection *mongo.Collection
}

func NewTareaStore(db *mongo.Database) *TareaStore {
	return &TareaStore{
		collection: db.Collection("tareas"),
	}
}

func (ts *TareaStore) Crear(tarea model.Tarea) error {
	_, err := ts.collection.InsertOne(context.TODO(), tarea)
	if err != nil {
		return err
	}
	return nil
}

func (ts *TareaStore) ObtenerTodos() ([]model.Tarea, error) {
	tareas := []model.Tarea{}
	valor, err := ts.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	err = valor.All(context.TODO(), &tareas)
	if err != nil {
		return nil, err
	}
	return tareas, nil

}

func (ts *TareaStore) ObtenerPorID(id string) (model.Tarea, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Tarea{}, err
	}
	resultado := ts.collection.FindOne(context.TODO(), bson.M{"_id": objectID})
	tarea := model.Tarea{}
	err = resultado.Decode(&tarea)
	if err != nil {
		return model.Tarea{}, err
	}
	return tarea, nil
}

func (ts *TareaStore) Actualizar(id string, tarea model.Tarea) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	objetoActualizar := bson.M{"_id": objectID}
	actualizacion := bson.M{"$set": tarea}
	_, err = ts.collection.UpdateOne(context.TODO(), objetoActualizar, actualizacion)
	if err != nil {
		return err
	}
	return nil
}

func (ts *TareaStore) Eliminar(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = ts.collection.DeleteOne(context.TODO(), bson.M{"_id": objectID})
	if err != nil {
		return err
	}
	return nil
}
