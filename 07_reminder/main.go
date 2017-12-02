package main

import "fmt"

func main(){

	for z:=0; z<1000; z++{
		if z % 2 == 1{											// the % check that the division has the rest or not
			fmt.Println("Bum")
		} else { fmt.Println("Bam") }
	}

}