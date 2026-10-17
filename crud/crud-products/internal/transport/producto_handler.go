package transport

import (
	"crud-products/internal/model"
	"crud-products/internal/service"

	"github.com/gin-gonic/gin"
)

type ProductoHandler struct {
	service *service.ProductoService
}

func NewProductoHandler(srv *service.ProductoService) *ProductoHandler {
	return &ProductoHandler{
		service: srv,
	}
}

func (ph *ProductoHandler) Crear(g *gin.Context) {
	var producto model.Producto
	if err := g.ShouldBindJSON(&producto); err != nil {
		g.JSON(400, gin.H{"error": "DATOS INVALIDOS"})
		return
	}
	err := ph.service.Crear(producto)
	if err != nil {
		g.JSON(500, gin.H{"error": err.Error()})
		return
	}
	g.JSON(201, gin.H{"creada": "Tarea creada"})
}

func (ph *ProductoHandler) ObtenerTodos(g *gin.Context) {
	productos, err := ph.service.ObtenerTodos()
	if err != nil {
		g.JSON(500, gin.H{"error": err.Error()})
		return
	}
	g.JSON(200, productos)
}

func (ph *ProductoHandler) ObtenerPorID(g *gin.Context) {
	obtenerID := g.Param("id")
	productoID, err := ph.service.ObtenerPorID(obtenerID)
	if err != nil {
		g.JSON(404, gin.H{"error": "este valor no puede estar vacio"})
		return
	}
	g.JSON(200, productoID)
}

func (ph *ProductoHandler) Actualizar(g *gin.Context) {
	productoActualizar := g.Param("id")

	var producto model.Producto

	if err := g.ShouldBindJSON(&producto); err != nil {
		g.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := ph.service.Actualizar(productoActualizar, producto)
	if err != nil {
		g.JSON(500, gin.H{"error": err.Error()})
		return
	}
	g.JSON(200, gin.H{"mensaje": "La tarea esta actualizada correcta mente"})
}

func (ph *ProductoHandler) Eliminar(g *gin.Context) {

	id := g.Param("id")
	err := ph.service.Eliminar(id)
	if err != nil {
		g.JSON(500, gin.H{"error": err.Error()})
		return
	}
	g.JSON(200, gin.H{"mensaje": "El producto se elimino de manera correcta"})

}
