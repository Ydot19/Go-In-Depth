// Echo2 prints its command line arguments
package main

import (
	"fmt"
	"os"
)

func main()  {
	s, sep := "", ""
	for _, arg := range os.Args[1:]{
		// note first element of Args is the actual command not the arguments
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}

/**
The following are equivalent
s := ""
var s string
var s = ""
var s string = ""


// Only the first two are valid
 */


