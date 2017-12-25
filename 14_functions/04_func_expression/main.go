package main

import "fmt"

func main(){

	print_func:=func(){								//declaring function like this is the only way to have a func inside another function

		fmt.Println("Hey There!")

	}

	print_func()
	fmt.Printf("%T\n",print_func)			//allow us to print the type of the var passed (in this case is func() type)
	
}
