package main

import "fmt"

func main(){

	i := 0

	for {
		fmt.Println(i)				// it will print every value from 0 without stop, till we kill the process
		i++
	}

}