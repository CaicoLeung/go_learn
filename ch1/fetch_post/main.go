package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type PostBody struct {
	Name string
	Age  int8
}

func main() {
	for _, url := range os.Args[1:] {
		params := PostBody{Name: "caico", Age: 18}
		jsonb, _ := json.Marshal(params)
		res, err := http.Post(url, "application/json", bytes.NewReader(jsonb))
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "fetch_post: %v\n", err)
			os.Exit(1)
		}
		body, err := ioutil.ReadAll(res.Body)
		_ = res.Body.Close()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stdout, "fetch_post: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", body)
	}
}
