package main

import (
	"fmt"
	"sync"
)

var counter int = 0
var wg sync.WaitGroup

func main() {
	var mutex sync.Mutex

	for i := 1; i < 5; i++ {
		go work(i, &mutex)
	}

	wg.Wait()
	fmt.Println("The End")
}

func work(number int, mutex *sync.Mutex) {
	defer wg.Done()
	wg.Add(1)
	mutex.Lock()
	counter = 0
	for k := 1; k <= 5; k++ {
		counter++
		fmt.Println("Goroutine", number, "-", counter)
	}
	mutex.Unlock()
}
