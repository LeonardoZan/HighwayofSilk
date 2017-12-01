package _1_memory_address

import "fmt"

func main(){

	a := 17

	fmt.Println("a - ",a)													// a -> 17
	fmt.Println("a memory address - ", &a)									// a -> 0xc420014068

	var b *int = &a																// giving b the int pointer type and value of &a address

	fmt.Println(b)		// 0xc420014068
	fmt.Println(*b)		// 17													// dereferencing the b var from the value of the address

}