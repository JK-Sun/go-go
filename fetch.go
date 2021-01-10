package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	args := os.Args
	fmt.Printf("args: %v\n", args)                    // [./fetch, a, b, c]
	fmt.Printf("args 1 to len(args): %v\n", args[1:]) // [a, b, c]

	for _, url := range args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		body, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading: %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", body)
		fmt.Printf("%s\n", resp.Status)
	}
}
