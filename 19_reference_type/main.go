package main

import "fmt"

func main(){

	m := make([]string, 1, 25)
	fmt.Println(m)						// []

	changeMe(m)

}


func changeMe(z []string){

	z[0]= "Leo"
	fmt.Println(z)						//[Leo]

}