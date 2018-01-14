package main

import "fmt"

func main(){

	var x [17]string 								// if between [] there isn't a number than the type is a slice
													// slice is dynamic, that means we can add as much value as we like
	fmt.Println(x)

	for i:=68;i<85;i++{
		x[-68+i] = string(i)
	}

	fmt.Println(x)

}
