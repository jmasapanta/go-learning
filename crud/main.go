package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Cargar variables del .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando .env")
	}

	// Leer la URI de MongoDB del .env
	uri := os.Getenv("MONGO_URI")
	dbName := os.Getenv("DB_NAME")

	// Conectar a MongoDB
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Error conectando a MongoDB:", err)
	}
	defer client.Disconnect(context.TODO())

	// Verificar conexión
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("No se pudo conectar a MongoDB:", err)
	}

	fmt.Println("Conectado a MongoDB Atlas:", dbName)
}
