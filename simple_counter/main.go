package main

import (
	"bufio" // read text
	"flag"
	"fmt" // printing stuff
	"io"  // read from io
	"os"  // use os resources?
)

func count(r io.Reader, countLines bool, countBytes bool) int {
    var scanner *bufio.Scanner

    if countLines {
        scanner = bufio.NewScanner(r)
    } else {
        scanner = bufio.NewScanner(r)
        scanner.Split(bufio.ScanWords)
    }

    if countBytes {
        scanner = bufio.NewScanner(r)
        scanner.Split(bufio.ScanBytes)
    }

    wc := 0

    for scanner.Scan() {
        wc++
    }

    return wc
}


func main() {
	lines := flag.Bool("l", false, "Count lines")
	line_bytes := flag.Bool("b", false, "count bytes")
	flag.Parse()
	fmt.Println(count(os.Stdin, *lines, *line_bytes))
}
