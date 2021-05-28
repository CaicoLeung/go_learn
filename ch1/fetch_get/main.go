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
		// 如果输入的url参数没有 https:// 前缀的话，为这个url加上该前缀
		if !strings.HasPrefix(url, "https://") && !strings.HasPrefix(url, "http://") {
			url = "https://" + url
		}
		fmt.Printf("url: %s\n", url)
		res, err := http.Get(url)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "fetch_post: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Http Status Code: %s\n", res.Status)
		_, err = io.Copy(os.Stdout, res.Body)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stdout, "fetch_post: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		_ = res.Body.Close()
	}
}
