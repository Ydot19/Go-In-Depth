package main

/**
Fetch prints the content found at a URL using command line arguments'

Requirement: Use io.Copy(dist, src) instead of ioutil.ReadAll to hold the
entire buffer
 */

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main(){
	fetchWithIOCopy()
}

func fetchWithIOCopy(){
	for _, url := range os.Args[1:]{
		resp, err := http.Get(url)
		if err != nil {
			_, er := fmt.Fprintf(os.Stderr, "Fetch Error: %s\n", err)
			if er != nil {
				fmt.Printf("Error in Fetch. Failed to Print Stderr: %s", er)
			}
			os.Exit(1)  // Exit code: 1
		}

		_, er := io.Copy(os.Stdout, resp.Body)

		if er != nil {
			fmt.Printf("Error reading response body: %s", err)
		}

		err = resp.Body.Close()

		if err != nil {
			fmt.Printf("Failed to close response body: %s", er)
		}

	}

}
