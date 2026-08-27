package main

import "fmt"

type Cuadrado struct {
	Base, Altura float64
}

func (c Cuadrado) Area() float64 {
	return c.Base * c.Altura
}

func str() {
	cuad := Cuadrado{
		Base:   5,
		Altura: 10,
	}
	fmt.Println(cuad.Area())
}
