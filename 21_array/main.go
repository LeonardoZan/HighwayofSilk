package main

import "fmt"

func main(){

	var x [17]string

	fmt.Println(x)

	for i:=68;i<85;i++{
		x[-68+i] = string(i)
	}

	fmt.Println(x)

}
