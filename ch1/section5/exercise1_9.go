package main

/**
Fetch prints the content found at a URL using command line arguments

Requirement: Also show the response status
 */

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main(){
	fetchWithResponseStatus()
}

func fetchWithResponseStatus(){
	for _, url := range os.Args[1:]{
		resp, err := http.Get(url)
		if err != nil {
			_, er := fmt.Fprintf(os.Stderr, "Fetch Error: %s\n", err)
			if er != nil {
				fmt.Printf("Error in Fetch. Failed to Print Stderr: %s", er)
			}
			os.Exit(1)  // Exit code: 1
		}

		b, er := ioutil.ReadAll(resp.Body)

		if er != nil {
			fmt.Printf("Error reading response body: %s", err)
		}

		err = resp.Body.Close()

		if err != nil {
			fmt.Printf("Failed to close response body: %s", er)
		}

		fmt.Printf("%s\n\n", b)
		fmt.Printf("HTTP STATUS CODE:\t %d\n", resp.StatusCode)
	}

}
