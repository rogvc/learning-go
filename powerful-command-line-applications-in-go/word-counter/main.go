package main

import (
	"bufio" // Buffered IO
	"flag"  // Used to read CLI flags
	"fmt"   // Used to output to stdout
	"io"    // Used to get a Reader interface
	"os"    // Used to retrieve input from stdin
)

func main() {
	ls := flag.Bool("l", false, "Count lines")
	bs := flag.Bool("b", false, "Count bytes")
	flag.Parse()

	fmt.Println(count(os.Stdin, *ls, *bs))
}

func count(r io.Reader, ls bool, bs bool) int {
	s := bufio.NewScanner(r)

	if bs {
		s.Split(bufio.ScanBytes)
		fmt.Println("looking for bytes")
	} else if ls {
		s.Split(bufio.ScanLines)
		fmt.Println("looking for lines")
	} else {
		s.Split(bufio.ScanWords)
		fmt.Println("looking for words")
	}

	var c int
	for s.Scan() {
		c++
	}

	return c
}
