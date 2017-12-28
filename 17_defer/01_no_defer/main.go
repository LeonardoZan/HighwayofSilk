package main

import (
	"fmt"
)

func hello(){fmt.Print("Hello")}
func world(){fmt.Println("World")}

func main(){

	world()
	hello()

}													// print out will be messy and not good as we desire
