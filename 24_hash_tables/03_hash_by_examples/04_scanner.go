package main

import (

	"fmt"
	"bufio"

	"strings"
	"os"
)

func main(){

	const input = "Once more into the Fray... Into the last good fight I'll ever know. Live and Die on this day... Live and die on this day..."
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan(){

		fmt.Println(scanner.Text())

	}

	if err := scanner.Err(); err != nil {

		fmt.Fprintln(os.Stderr, "reading input", err)

	}

}