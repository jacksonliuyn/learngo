package algorithm

import (
	"fmt"
	"sync"
)

func tworotine() {
	wg := sync.WaitGroup{}
	ch1, ch2 := make(chan struct{}, 1), make(chan struct{}, 1)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i += 2 {
			ch1 <- struct{}{}
			fmt.Println(i)
			<-ch2
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i < 100; i += 2 {
			<-ch1
			fmt.Println(i)
			ch2 <- struct{}{}

		}
	}()

	wg.Wait()
}
