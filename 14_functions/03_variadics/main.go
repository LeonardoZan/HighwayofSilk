package main

import "fmt"

func main(){

	num := average( 22, 17, 29, 7, 89, 353, 11, 24, 9)
	fmt.Println(num)

}

func average(sf ...float64) float64{			// 3 dots make the function variadics, that means that I can add as much parameters as much I need to.

	var total float64							// declaring the var like this it will set it automatically to zero, it's an optimal way to declare vars

	for _, v:= range sf {						// range allow us to loop many times as the number of params passed to the function

		total+=v								// attention to the fact that var v assumes the value of each parameter passed, until the end of them
												// every loop assign to v the var of a parameter and check if there are more, if yes the loop continue
												// if negative, the loop stop because there are no more parameters passed

	}

	return total/float64(len(sf))				// len(sf) check the length of the list, so we can divide the total by the number of params passed and obtain the average
}
