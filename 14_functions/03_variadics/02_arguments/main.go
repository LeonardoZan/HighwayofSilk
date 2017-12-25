package main

import "fmt"

func main(){

	data := []float64{22, 17, 29, 7, 89, 353, 11, 24, 9}
	num := average(data...)												//passing arguments 1 by 1, important to remember that in this case "..." goes after the var name to pass
	fmt.Println(num)

}

func average(sf ...float64) float64{

	var total float64

	for _, v:= range sf {

		total+=v



	}

	return total/float64(len(sf))
}
