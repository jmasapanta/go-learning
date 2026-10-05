package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Producto struct{
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Nombre string `json:"nombre" bson:"nombre"`
	Precio float64 `json:"precio" bson:"precio"`
	Stock int `json:"stock" bson:"stock"`
	Categoria string  `json:"categoria" bson:"categoria"`
	Disponible bool `json:"disponible" bson:"disponible"`
}