package main

import "fmt"

func main(){

	a := 17

	fmt.Println("a - ",a)
	fmt.Println("a memory address - ", &a)				// memory address location of the const a
	fmt.Printf("%d \n", &a)							// decimal output memory address

}
