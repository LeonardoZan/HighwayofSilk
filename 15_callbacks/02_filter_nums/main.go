package main

import "fmt"

func main(){

	xs:= filter([]int{12,354,124,77}, func(n int) bool {

		return n>100

	})

	fmt.Println(xs)

}


func filter(numbers []int, callback func(int) bool) []int{

	var xs []int
	for _, n:=range numbers{

		if callback(n){

			xs =append(xs,n)

		}

	}
	return xs
}