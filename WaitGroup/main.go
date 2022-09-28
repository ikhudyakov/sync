package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	count := 0

	wgroup(&wg, &count)
	wgroup(&wg, &count)
	wgroup(&wg, &count)

	wg.Wait()
	fmt.Println("All goroutines complete.")
}

func wgroup(wg *sync.WaitGroup, count *int) {
	wg.Add(1)
	go func() {
		*count++
		defer wg.Done()
		fmt.Printf("%dst goroutine\n", *count)
		time.Sleep(1000 * time.Millisecond)
	}()
}
