package main

/**
Prints the count and text of lines that appear more than once in the input.
It reads from stdin or from a list of name files
 */

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	counts := make(map[string]int)
	files := os.Args[1:] // Note: Arg 0 is the executable/command path
	if len(files) == 0 {
		// Note: os.Stdin follows the File Struct in go
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				_, err = fmt.Fprintf(os.Stderr, "dup2 error: %v\n", err)
				if err != nil {
					fmt.Println("Os.Open Error. Failed to Print")
				}
				continue
			}
			countLines(f, counts)
			err = f.Close()
			if err != nil {
				fmt.Printf("Failed to clse file: %s", arg)
			}
		}
	}

	for word, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, word)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		for _, word := range strings.Fields(scanner.Text()){
			counts[word]++
		}
	}
}