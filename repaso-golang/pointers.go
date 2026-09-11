package main

import "fmt"

func pointersd() {
	x := 10
	p := &x
	fmt.Println("Valor de x:", x)
	fmt.Println("Direccion de x:", p)
	fmt.Println("Valor apuntado por p:", *p)
}
/*
package main

import "fmt"

func incrementarPorValor(x int) {
	x = x + 1
	fmt.Println("Dentro de la funcion:", x)
}
func imcrementaPorPointers(x *int) {
	*x = *x + 1
	fmt.Println("Dentro de la funcion:", *x)
}

func main() {
 
	num := 10

	incrementarPorValor(num)
	fmt.Println("Despues de incrementar por valor:", num)

	imcrementaPorPointers(&num)
	fmt.Println("Despues de incrementar por pointer:", num)
}
*/
