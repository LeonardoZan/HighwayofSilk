package main

import "fmt"

func main(){

	myGreetings := map[string]string{

		"Oh" :		"Yes dude, Oh",
		"Ah" :		"That's what I'm talking about! Ah",

	}

	myGreetings["Leo"] = "How you doing?"
	myGreetings["Max"] = "Hey there!"



	fmt.Println(myGreetings)
}
