package main

import (
	"fmt"
	"os"
)

/**
Modify the echo problem to also print os.Args prints the command that invoke it
 */
func main()  {
	var s, sep string

	for _, arg := range os.Args{
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}