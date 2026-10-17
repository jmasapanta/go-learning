package main

import (
	"context"
	"crud-products/internal/service"
	"crud-products/internal/transport"
	"fmt"
	"log"
	"os"

	// Faltan estos imports

	"crud-products/internal/store"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando env")
	}
	url := os.Getenv("MONGO_URI")
	db := os.Getenv("DB_NAME")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal("error conectando a mongo", err)
	}

	defer client.Disconnect(context.TODO())

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("no se pudo conectar", err)
	}
	fmt.Println("conectadoa mongo db", db)

	database := client.Database(db)
	ProductStore := store.NewProductoStore(database)
	ProductService := service.NewProductoService(ProductStore)
	ProductHandler := transport.NewProductoHandler(ProductService)
	r := gin.Default()

	r.POST("/productos", ProductHandler.Crear)
	r.GET("/productos", ProductHandler.ObtenerTodos)
	r.GET("/productos/:id", ProductHandler.ObtenerPorID)
	r.PUT("/productos/:id", ProductHandler.Actualizar)
	r.DELETE("/productos/:id", ProductHandler.Eliminar)

	r.Run(":8080")

}
