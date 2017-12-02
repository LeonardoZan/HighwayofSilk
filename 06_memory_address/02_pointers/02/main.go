package main

import (
	"fmt"
)

func zero(z *int){

	fmt.Printf("%p\n",&z)					// address in function zero
	*z = 0

}


func main(){

	x := 5
	fmt.Printf("%p\n", &x)					// address in main

	zero(&x)
	fmt.Println(x)									// x is 0

}