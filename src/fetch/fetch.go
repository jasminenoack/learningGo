// Package fetch prints the content found at each specifiedc URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// b, err := ioutil.ReadAll(resp.Body)
		// resp.Body.Close()
		// if err != nil {
		// 	fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		// 	os.Exit(1)
		// }
		// fmt.Printf("%s", b)
		fmt.Println("")
		fmt.Println(resp.Status)
		fmt.Println("")
		_, err = io.Copy(os.Stdout, resp.Body)
		fmt.Println("")
		fmt.Println("")
		if err != nil {
			fmt.Fprintf(os.Stderr, "copy: %v", err)
		}
	}
}
