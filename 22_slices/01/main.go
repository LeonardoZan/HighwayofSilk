package _1

import "fmt"

func main(){

	my_slice := []int{1,2,3,4,5,6,25,145,658,63,32}					// if between [] there isn't a number than the type is a slice
	fmt.Printf("%T\n", my_slice)
	fmt.Println(my_slice)
	fmt.Println(len(my_slice))


}
