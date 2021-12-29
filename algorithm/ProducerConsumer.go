package algorithm

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

//func main() {
//	//先写个同步的
//	ch := make(chan int)
//	wg.Add(2)
//	go producer(ch)
//	go consumer(ch)
//	wg.Wait()
//}

func consumer(ch chan int) {
	for {
		num, ok := <-ch
		if ok {
			fmt.Println(num)
		} else {
			break
		}
	}
	wg.Done()
}
func producer(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	// close(ch)
	wg.Done()
}
