package _1_simple_example

import "fmt"

func main(){

	yesman:="String to print"
	print_string_func(yesman)				// parameters can be as a var (in this case: yesman)or as a string (print_string_func("Hey There!")

}

func print_string_func(name string){		// defining what the function do

	fmt.Println(name)
}