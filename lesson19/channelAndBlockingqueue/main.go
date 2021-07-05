package main

import (
	"fmt"
	"time"
)

func main() {
	ch := Request("https:/github.com")
	select {
	case r := <- ch:
		fmt.Println(r)
	}
}

func Request(url string) <-chan string  {
	ch := make(chan string)
	go func() {
		time.Sleep(time.Second)
		ch <- fmt.Sprintf("url = %s, res = %s", url, "ok")
	}()
	return ch
}