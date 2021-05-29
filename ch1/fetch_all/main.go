package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string) // make函数创建了一个传递string类型参数的channel
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // 从ch channel receive
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

/*
当一个goroutine尝试在一个channel上做send或者receive操作时，这个goroutine会阻塞在调用处，直到另一个goroutine从这个channel里接收或者写入值，
这样两个goroutine才会继续执行channel操作之后的逻辑。
在这个例子中，每一个fetch函数在执行时都会往channel里发送一个值（ch <- expression），主函数负责接收这些值（<-ch）。
这个程序中我们用main函数来接收所有fetch函数传回的字符串，可以避免在goroutine异步执行还没有完成时main函数提前退出。
*/
func fetch(url string, ch chan<- string) {
	start := time.Now()
	res, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send至ch channel
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, res.Body) // 可以把ioutil.Discard这个变量看作一个垃圾桶，可以向里面写一些不需要的数据
	if err != nil {
		ch <- fmt.Sprintf("While reading %s: %v\n", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
