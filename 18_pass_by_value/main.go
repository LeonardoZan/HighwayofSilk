package main

import "fmt"

func main(){

	age := 18
	fmt.Println(&age)			// 0xc420014068

	changeMe(&age)

	fmt.Println(&age)			// 0xc420014068
	fmt.Println(age)			// 22

}

func changeMe(z *int){

	fmt.Println(z)
	fmt.Println(*z)				// (showing 18) dereference, showing the value contained in the address z (address of var age)

	*z = 22

	fmt.Println(z)
	fmt.Println(*z)

}
