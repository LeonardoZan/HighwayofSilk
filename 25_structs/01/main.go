package main

import "fmt"

func main(){

	p1 := person{"Jack", "Sparrow", 30}
	p2 := person{"Ragnar", "Lothbrok", 40}

	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)

}

type person struct{

	first string
	last string
	age int
}