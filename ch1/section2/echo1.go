// Creating the echo function in the terminal
package main

import (
	"fmt"
	"os"
)

func main()  {
	var s, sep string

	for i := 1; i< len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)
}

/**
Run with the following

go run echo1.go Running Echo Function

or

go build echo1.go

./echo1 Running Echo Function

both prints out the following...
Running Echo Function
 */
