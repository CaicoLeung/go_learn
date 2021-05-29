package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"time"
)

func main() {
	total := 0 // 请求的总次数
	start := time.Now()
	ch := make(chan map[string]string) // make函数创建了一个传递string类型参数的channel
	for _, url := range os.Args[1:] {
		fmt.Printf("Start fetch %v\n", url)
		go fetch(url, ch) // start a goroutine
		go fetch(url, ch) // start a goroutine
		total += 2
	}
	result := make(map[string]string)
	var keys []string
	for i := 0; i < total; i++ {
		for k, v := range <-ch {
			result[k] = v
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Printf("%s: %s\n", key, result[key])
	}
	fmt.Printf("done, use %v seconds", time.Since(start).Seconds())
}

/*
当一个goroutine尝试在一个channel上做send或者receive操作时，这个goroutine会阻塞在调用处，直到另一个goroutine从这个channel里接收或者写入值，
这样两个goroutine才会继续执行channel操作之后的逻辑。
在这个例子中，每一个fetch函数在执行时都会往channel里发送一个值（ch <- expression），主函数负责接收这些值（<-ch）。
这个程序中我们用main函数来接收所有fetch函数传回的字符串，可以避免在goroutine异步执行还没有完成时main函数提前退出。
*/
func fetch(url string, ch chan<- map[string]string) {
	start := time.Now()
	res, err := http.Get(url)
	result := make(map[string]string)
	if err != nil {
		result[url] = fmt.Sprintf("http-get: %v\n", err) // send至ch channel
		ch <- result
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, res.Body) // 可以把ioutil.Discard这个变量看作一个垃圾桶，可以向里面写一些不需要的数据
	if err != nil {
		result[url] = fmt.Sprintf("While reading %s: %v\n", url, err)
		ch <- result
		return
	}
	secs := time.Since(start).Seconds()
	result[url] = fmt.Sprintf("%v %v %v Bytes %v", url, res.Status, nbytes, secs)
	ch <- result
}
