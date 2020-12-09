package main

import (
	"fmt"
	"os"
)

/**
Modify the echo program to print the index and value of each of its arguments, one per line
 */
func main()  {

	for  i :=1; i < len(os.Args); i++ {
		fmt.Println("Index: ", i, "- Argument: ", os.Args[i])
	}
}