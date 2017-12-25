package main

import "fmt"

func main(){

	fmt.Println(double_print_func("Highway","ofSilk"))

}

func double_print_func(first_name, last_name string) (string ,string) {				// outside () we declare the type of the return we want (in this case string)

	return fmt.Sprint(first_name, last_name), fmt.Sprint(last_name, first_name)		//the func Sprint print the values to a string not to the monitor, is just kind of concatenating values

}
