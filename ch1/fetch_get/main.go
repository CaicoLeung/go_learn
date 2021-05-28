package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		res, err := http.Get(url)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "fetch_post: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, res.Body)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stdout, "fetch_post: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		_ = res.Body.Close()
	}
}
