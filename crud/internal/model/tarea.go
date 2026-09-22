package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Tarea struct {
    ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
    Nombre       string             `json:"nombre" bson:"nombre"`
    Descripcion  string             `json:"descripcion" bson:"descripcion"`
    Prioridad    string             `json:"prioridad" bson:"prioridad"`
    Materia      string             `json:"materia" bson:"materia"`
    CantidadHoras int              `json:"cantidad_horas" bson:"cantidad_horas"`
    Estado       bool               `json:"estado" bson:"estado"`
    FechaEntrega string             `json:"fecha_entrega" bson:"fecha_entrega"`
}