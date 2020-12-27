package main

/**
Simplification of dup2.go without the functionality for observing os.Stdin if no
file arguments are provided
 */

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	// note Arg at position 0 represents the command call and not the
	// actual arguments
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			_, err = fmt.Fprintf(os.Stderr, "dup3 error: %v\n", err)
			if err != nil {
				fmt.Printf("fmt.Fprint Error %v\n", err)
			}
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			for _, word := range strings.Fields(line){
				counts[word]++
			}
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
