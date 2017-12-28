package main

import (
	"fmt"
)

func hello(){fmt.Print("Hello ")}
func world(){fmt.Println("World")}

func main(){

	defer world()									//defer is used to make the execution of a func just a moment before going out of the main, so in the end
	hello()

}