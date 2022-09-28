package main

import (
	"fmt"
	"time"
)

func main() {
	var ch chan int = make(chan int, 3)
	var v int = 0

	go func() {
		for {
			fmt.Println(<-ch)
		}
	}()

	for {
		v = v + 1
		ch <- v
		time.Sleep(time.Second)
	}
}
