package main

import "fmt"

func main(){

	var myGreetings = make(map[string]string)
	myGreetings["Leo"] = "How you doing?"
	myGreetings["Max"] = "Hey there!"

	fmt.Println(myGreetings)
}
