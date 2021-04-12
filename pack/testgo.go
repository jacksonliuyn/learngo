package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var lock sync.Mutex

func main() {
	var x int64
	wg.Add(2)

	go func() {
		lock.Lock()
		for i := 0; i < 5000; i++ {

			x = x + 1

		}
		lock.Unlock()
		wg.Done()
	}()
	go func() {
		lock.Lock()
		for i := 0; i < 5000; i++ {
			x = x + 1
		}
		lock.Unlock()
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(x)
}
