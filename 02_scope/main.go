package main

import "fmt"

var x int = 42						// type can be omitted
									// the way we have declared the var can be defined as a package level scope,
									// meaning that we have access to it in from all over the package

func main()  {

	fmt.Println(x)					// OUTPUT ==> 	42
	foo()							// OUTPUT ==>	42

	y := 7							// this var is only accessible within this function because is declared "locally" inside of it

	fmt.Println(y)

}

func foo()  {

	fmt.Println(x)
//	fmt.Println(y)					// the print here won't work because of the declaration of the var inside a function different from
									// the one we're calling the function from
}
