package main

import (
	"errors"
	"fmt"
)

func pedirInformacion() (int, string, float64) {
	var mesa int
	var nombrePlato string
	var precio float64
	fmt.Print("Ingrese la mesa: ")
	fmt.Scan(&mesa)
	fmt.Print("Nombre del plato: ")
	fmt.Scan(&nombrePlato)
	fmt.Print("Precio del plato: ")
	fmt.Scan(&precio)
	return mesa, nombrePlato, precio
}

func calcularTotal(precios []float64, estados []string) (float64, error) {
	totalVentas := 0.0
	if len(precios) == 0 {
		return 0, errors.New("No precios")
	}
	for i := 0; i < len(precios); i++ {
		if estados[i] == "entregado" {
			totalVentas += precios[i]
		}
	}
	return totalVentas, nil
}

func main() {
	fmt.Println("=== SISTEMA PEDIDOS ===")
	mesas := []int{}
	platos := []string{}
	precios := []float64{}
	estados := []string{}

	for {
		fmt.Println("1. Agregar pedido")
		fmt.Println("2. Ver todos los pedidos")
		fmt.Println("3. Buscar pedido por mesa")
		fmt.Println("4. Ver total recaudado")
		fmt.Println("5. Ver pedidos pendientes")
		fmt.Println("6. Salir")
		fmt.Print("\nElige una opcion: ")

		var opcion int
		fmt.Scan(&opcion)

		switch opcion {
		case 1:
			estadoPedido := "Pendiente"
			mesa, nombrePlato, precio := pedirInformacion()
			mesas = append(mesas, mesa)
			platos = append(platos, nombrePlato)
			precios = append(precios, precio)
			estados = append(estados, estadoPedido)
			fmt.Printf("Pedido entregado en mesa : %d | Plato: %s | Precio: %.2f\n", mesa, nombrePlato, precio)

		case 2:
			if len(mesas) == 0 {
				fmt.Println("No existen pedidos")
			} else {
				for i := 0; i < len(mesas); i++ {
					fmt.Printf("Mesa: %d | Plato: %s | Precio: %.2f | Estado: %s\n", mesas[i], platos[i], precios[i], estados[i])
				}
			}

		case 3:
			var pedirMesa int
			fmt.Scan(&pedirMesa)
			encontrado := false
			for i := 0; i < len(mesas); i++ {
				if mesas[i] == pedirMesa {
					fmt.Printf("mesas: %d | platos:%s | precios: %.2f | estados: %s ", mesas[i], platos[i], precios[i], estados[i])
					encontrado = true
				}
			}
			if !encontrado {
				fmt.Print("Mesa no encontrada")
			}

		case 4:
			total, err := calcularTotal(precios, estados)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("Total recaudado:  %.2f\n", total)
		
		case 5:
			pedidosPendientes:=[]string{}
			for i := 0; i < len(estados); i++ {
				if estados[i] == "Pendiente" {
					pedidosPendientes = append(pedidosPendientes, fmt.Sprintf("Mesa: %d | Plato: %s | Precio: %.2f", mesas[i], platos[i], precios[i]))
				}
			}
			if len(pedidosPendientes) == 0 {
				fmt.Println("No hay pedidos pendientes")
			} else {
				fmt.Println("Pedidos pendientes:")
				for _, pedido := range pedidosPendientes {
					fmt.Println(pedido)
				}
			}


		case 6:
			fmt.Println("Hasta luego!")
			return
		default:
			fmt.Println("Opcion invalida")
		}

	}

}
