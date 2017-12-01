package main

import (

	"net/http"
	"io/ioutil"
	"fmt"
)
func main() {

	res, _ := http.Get("http://testsite.test/")						//take the page
	page, _ := ioutil.ReadAll(res.Body)									// read all the body and put it in a var
	res.Body.Close()													// we close the response body
	fmt.Printf("%s", page)										// we print all

	// the blank identifier _ throw away any error that may came
}