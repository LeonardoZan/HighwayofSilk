package main

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	)

func main(){

	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")					// gaining the .txt file from the URL
	if err!=nil{
		log.Fatalln(err)
	}

	bs, _ := ioutil.ReadAll(res.Body)																						// getting slice of bytes in the var bs
	str:=string(bs)																											// converting bs slice of bytes into a string
	fmt.Println(str)

}
