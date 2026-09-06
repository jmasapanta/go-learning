package main

import "fmt"

type Hablador interface {
	Hablar() string
}
type Perro struct{}

func (p Perro) Hablar() string {
	return "Guau"
}

func inter() {
	var p Hablador = Perro{}
	fmt.Println(p.Hablar())
}
