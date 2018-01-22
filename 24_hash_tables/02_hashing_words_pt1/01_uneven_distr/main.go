package main

import (
	"fmt"
	"net/http"
	"log"
	"bufio"
)

func main() {

	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt") 							// gaining the moby dick book
	if err != nil {
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(res.Body) 																		// scanning the page
	defer res.Body.Close()

	scanner.Split(bufio.ScanWords)																				// setting the split function for the scan operation
	buckets := make([]int, 12)																					// creating slice to hold count

	for scanner.Scan(){

		n := HashBuckets(scanner.Text())
		buckets[n]++

	}

	fmt.Println(buckets[65:122])

}

func HashBuckets(word string) int{

	return int(word[0])																							// takes the first letter of the word and convert it into a number

}
