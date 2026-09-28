package transport

import (
	"crud/internal/model"
	"crud/internal/service"

	"github.com/gin-gonic/gin"
)

type TareaHandler struct {
	service *service.TareaService
}

func NewTareaHandler(srv *service.TareaService) *TareaHandler {
	return &TareaHandler{
		service: srv,
	}
}

func (th *TareaHandler) Crear(c *gin.Context) {
	var tarea model.Tarea
	if err := c.ShouldBindJSON(&tarea); err != nil {
		c.JSON(400, gin.H{"error": "datos invalidos"})
		return
	}

	err := th.service.Crear(tarea)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"mensaje": "tarea creada"})
}

func (th *TareaHandler) ObtenerTodos(c *gin.Context) {
	tareas, err := th.service.ObtenerTodos()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, tareas)
}
func (th *TareaHandler) ObtenerPorID(c *gin.Context) {
	id := c.Param("id")
	tarea, err := th.service.ObtenerPorID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Datos incorrectos"})
	}
	c.JSON(200, tarea)
}

func (th *TareaHandler) Actualizar(c *gin.Context) {
	id := c.Param("id")
	var tarea model.Tarea
	if err := c.ShouldBindJSON(&tarea); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := th.service.Actualizar(id, tarea)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"mensaje": "tarea actualizada"})
}

func (th *TareaHandler) Eliminar(c *gin.Context) {
	id := c.Param("id")
	err := th.service.Eliminar(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"mensaje": "Tarea eliminada"})
}
