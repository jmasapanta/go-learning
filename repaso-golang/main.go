package main

import "fmt"

func main() {
	slice:=[]int{1,3,5}
	slice = append(slice, 2,4)
	fmt.Println(slice)
}
