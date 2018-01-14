package main

import "fmt"

func main(){

	my_slice := make([]int,0,10)						// appending more vales to the slices than the size of the underlying array is cost performance
														// because every time we exceed the size of the array there's the creation of a new one double-size n
	fmt.Println("- - - - - - - - - - - - -")
	fmt.Println(my_slice)
	fmt.Println(len(my_slice))
	fmt.Println(cap(my_slice))
	fmt.Println("- - - - - - - - - - - - -")

	for i := 0; i < 80; i++ {

		my_slice = append(my_slice, i)
		fmt.Println("len: ", len(my_slice), "capacity: ", cap(my_slice), "value: ", my_slice[i])

	}

}
