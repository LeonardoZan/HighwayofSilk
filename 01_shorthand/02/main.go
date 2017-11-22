package main

import "fmt"

func main()  {

	a := 10
	b := "golang"
	c := 4.17
	d := true

	fmt.Printf("%T \n", a)		// by using %T will print the Go representation of the var
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)


}