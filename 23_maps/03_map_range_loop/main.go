package main

import "fmt"

func main(){

	myGreeting := map[int]string{

		0:	"Good Morning",
		1:	"Bonjour",
		2:	"Buen Dias",
		3:	"Buongiorno",

	}

	for key, val := range myGreeting{		// everything is "iterable" can be used with range

		fmt.Println(key," - ", val)

	}

}