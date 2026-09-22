package service

import (
	"errors"
	"crud/internal/model"
	"crud/internal/store"
)

type TareaService struct {
	store *store.TareaStore
}

func NewTareaService(trs *store.TareaStore) *TareaService {
	return &TareaService{
		store: trs,
	}
}
func (trs *TareaService) Crear(tarea model.Tarea) error {
	if tarea.Nombre == "" {
		return errors.New("el nombre de la tarea no puede estar vacío")
	}
	err:=trs.store.Crear(tarea)
	if err !=nil{
		return err
	}
	return nil
}
func (trs *TareaService) ObtenerTodos()([]model.Tarea, error){
	return trs.store.ObtenerTodos()
}

func (trs *TareaService) ObtenerPorID(id string)(model.Tarea,error){
	return trs.store.ObtenerPorID(id)
}

func (trs *TareaService) Actualizar(id string, tarea model.Tarea)error{
	if tarea.Nombre==""{
		return  errors.New("Nombre no puede estar vacio")
	}
	return trs.store.Actualizar(id,tarea)
}
func (trs *TareaService)Eliminar(id string) error{
	return trs.store.Eliminar(id)
}