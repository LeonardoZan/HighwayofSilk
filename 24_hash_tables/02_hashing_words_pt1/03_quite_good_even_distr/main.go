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

		n := HashBuckets(scanner.Text(),12)
		buckets[n]++

	}

	fmt.Println(buckets)

}

func HashBuckets(word string, buckets int) int{

	var sum int
	for _, v := range word{																						// for each word I'm going to hash

		sum += int(v)																								// I add the value of each letter

	}
	return sum % buckets																						// I, in the end return the final sum/number_of_buckets

}
