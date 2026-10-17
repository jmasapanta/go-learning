package service

import (
	"crud-products/internal/model"
	"crud-products/internal/store"
	"errors"
)

type ProductoService struct {
	store *store.ProductoStore
}

func NewProductoService(prd *store.ProductoStore) *ProductoService {
	return &ProductoService{
		store: prd,
	}
}

func (pr *ProductoService) Crear(producto model.Producto) error {
	if producto.Nombre == "" {
		return errors.New("El nombre no puede estar vacío")
	}
	err := pr.store.Crear(producto)
	if err != nil {
		return err
	}
	return nil
}

func (pr *ProductoService) ObtenerTodos() ([]model.Producto, error) {
	return pr.store.ObtenerTodos()
}

func (pr *ProductoService) ObtenerPorID(id string) (model.Producto, error) {
	return pr.store.ObtenerPorID(id)
}
func (pr *ProductoService) Actualizar(id string, producto model.Producto) error {
	if producto.Nombre == "" {
		return errors.New("nombre vacio")
	}
	return pr.store.Actualizar(id, producto)
}
func (pr *ProductoService) Eliminar(id string) error {
	return pr.store.Eliminar(id)
}
