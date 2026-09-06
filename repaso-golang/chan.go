package main

import "fmt"

func canal() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() { ch1 <- 2 }()
	go func() { ch2 <- 5 }()

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	}
}
