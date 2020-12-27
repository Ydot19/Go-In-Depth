package main
/**
Prints the text of each line that appears more than once in the standard output
Preceding by its count
**/

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	/**
	Holds a set of key/value pairs and provides constant time operations to store,
	retrieve or test for an item in the set
	 */
	counts := make(map[string]int)
	/**
	Gets the os standard inputs as an input
	 */
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		// zero value of int is 0
		counts[scanner.Text()]++
		/** Same as the following:
		line := scanner.Text()
		counts[line] = counts[line] + 1
		 */
		if scanner.Text() == ""{
			break
		}
	}

	// NOTE: program ignores errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}