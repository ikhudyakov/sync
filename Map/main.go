package main

import (
	"fmt"
	"sync"
)

func main() {
	var myMap sync.Map
	var wg sync.WaitGroup

	myMap.Store("qwerty", 1)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			myMap.Store("qwerty", 1)
		}()
	}
	wg.Wait()
	fmt.Println(myMap.Load("qwerty"))
}
