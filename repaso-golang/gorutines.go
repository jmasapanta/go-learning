package main

import (
	"fmt"
	"time"
)

func imprimirNumeros(n int) {
	for i := 0; i <= n; i++ {
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main2() {
	go imprimirNumeros(4)
	go imprimirNumeros(4)

	time.Sleep(time.Second)
}
