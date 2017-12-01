package main

import "fmt"

func main(){

	a := 17

	fmt.Println("a - ",a)													// a -> 17
	fmt.Println("a memory address - ", &a)									// a -> 0xc420014068

	var b *int = &a																// giving b the int pointer type and value of &a address

	fmt.Println(b)		// 0xc420014068
	fmt.Println(*b)		// 17													// dereferencing the b var from the value of the address

	*b = 42				// b says, "Now the value at this address is 42"
	fmt.Println(a)		// 42

}