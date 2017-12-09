package main

import "fmt"

func main() {

	switch "Test" {

	case "Test_1":
		fmt.Println("Ehy there! Test 1")

	case "Test_2":
		fmt.Println("Ehy there! Test 2")
	//	fallthrough									// this would mean that if the prev. case is true also this case is true as well

	case "Test_3":
		fmt.Println("Ehy there! Test 3")			// no break is needed

	}

}