package main

import "fmt"

func main(){

	fmt.Println(factorial(12))						// recursion is somehow "mirror to mirror", function calling itself
														// in this case, passing 12 we'll have (12*11*10...1) until we reach 0 --> we exit the func
}


func factorial(x int) int{

	if x==0 {

		return 1
	}
	return x * factorial(x-1)							// (12*11*10...1)
}
