package main

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
)

func main() {

	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt") 							// gaining the moby dick book
	if err != nil {
		log.Fatalln(err)
	}

	bs, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil{

		log.Fatal(err)

		}

	fmt.Printf("%s", bs)																					// printing the slice of bytes, bs, to string by using %s

}
