package main

import (
	"errors"
	"fmt"
	"math"
)

func sumar(a, b float64) float64 {
	return a + b
}
func restar(a, b float64) float64 {

	return a - b
}
func multiplicar(a, b float64) float64 {
	return a * b
}
func division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("no se puede dividir entre cero")
	}
	return a / b, nil
}
func potencia(base, exponente float64) float64 {
	return math.Pow(base, exponente)
}
func raizCuadrada(n float64) (float64, error) {
	if n < 0 {
		return 0, errors.New("no se puede calcular raiz de numero negativo")
	}
	return math.Sqrt(n), nil
}
func leerDosNumeros() (float64, float64) {
	var a, b float64
	fmt.Print("Primer numero: ")
	fmt.Scan(&a)
	fmt.Print("Segundo numero: ")
	fmt.Scan(&b)
	return a, b
}
func main() {
	fmt.Println("=== CALCULADORA GO ===")
	historial := []string{} // afuera del for
	for {
		fmt.Println("1. Suma")
		fmt.Println("2. Resta")
		fmt.Println("3. Multiplicacion")
		fmt.Println("4. Division")
		fmt.Println("5. Potencia")
		fmt.Println("6. Raiz cuadrada")
		fmt.Println("7. Ver historial")
		fmt.Println("8. Salir")
		fmt.Print("\nElige una opcion: ")

		var opcion int
		fmt.Scan(&opcion)

		switch opcion {
		case 1:
			a, b := leerDosNumeros()
			resultado := sumar(a, b)
			fmt.Printf("Resultado: %.2f\n", resultado)
			historial = append(historial, fmt.Sprintf("Suma: %.2f + %.2f = %.2f", a, b, resultado))
		case 2:
			a, b := leerDosNumeros()
			resultado := restar(a, b)
			fmt.Printf("Resultado: %.2f\n", resultado)
			historial = append(historial, fmt.Sprintf("Resta: %.2f - %.2f = %.2f", a, b, resultado))
		case 3:
			a, b := leerDosNumeros()
			resultado := multiplicar(a, b)
			fmt.Printf("Resltado: %.2f\n", resultado)
			historial = append(historial, fmt.Sprintf("Multiplicacion: %.2f * %.2f = %.2f", a, b, resultado))
		case 4:
			a, b := leerDosNumeros()
			resultado, err := division(a, b)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			fmt.Printf("Resultado: %.2f\n", resultado)
			historial = append(historial, fmt.Sprintf("Division: %.2f / %.2f = %.2f", a, b, resultado))
		case 5:
			base, exponente := leerDosNumeros()
			resultado := potencia(base, exponente)
			fmt.Printf("Resultado: %.2f\n", resultado)
			historial = append(historial, fmt.Sprintf("Potencia: %.2f ^ %.2f = %.2f", base, exponente, resultado))
		case 6:
			var n float64
			fmt.Print("Primer numero: ")
			fmt.Scan(&n)
			resultado, err := raizCuadrada(n)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			fmt.Printf("Resultado: %.2f\n", resultado)
			historial = append(historial, fmt.Sprintf("Raiz cuadrada: √%.2f = %.2f", n, resultado))
		case 7:
			if len(historial) == 0 {
				fmt.Println("No hay operaciones en el historial")
			} else {
				for i, operacion := range historial {
					fmt.Printf("%d. %s\n", i+1, operacion)
				}
			}
		case 8:
			fmt.Println("Hasta luego!")
			return
		default:
			fmt.Println("Opcion invalida")
		}
	}
}
