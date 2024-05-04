package main

import (
	"bufio" // read text
	"fmt"  // printing stuff 
	"io" // read from io	
	"os" // use os resources?
	//"strings"
)

func count(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	wc:= 0

	for scanner.Scan() {
		wc++
	}
	
	return wc;
}

func main() {
	fmt.Print(count(os.Stdin));	
}
