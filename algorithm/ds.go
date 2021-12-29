package algorithm

import "container/list"

type Stack struct {
	cache []int
}

func (sk *Stack) push(n int) {
	sk.cache = append(sk.cache, n)
}

func (sk Stack) length() int {
	return len(sk.cache)
}
func (sk *Stack) pop() int {
	if sk.length() == 0 {
		return 0
	}
	item := sk.cache[sk.length()-1]
	sk.cache = sk.cache[:len(sk.cache)-1]
	return item
}
func (sk Stack) isEmpty() bool {
	return len(sk.cache) == 0
}
func (sk Stack) peer() int {
	return sk.cache[sk.length()-1]
}
func Qs() {
	{
		// 初始化
		queue := list.New()
		stack := list.New()
		// 入队 	入栈
		queue.PushBack(123)
		stack.PushBack(123)
		// 出队 出栈 返回的数据是结构类型 Value 需要断言成相应的类型
		num1 := queue.Front()
		queue.Remove(num1)

		num2 := queue.Back()
		stack.Remove(num2)
	}
}

//package main
//
//import "fmt"
//
//func Producer(ch chan int) {
//　　for i := 1; i <= 10; i++ {
//　　　　ch <- i
//　　}
//　　close(ch)
//}
//
//func Consumer(id int, ch chan int, done chan bool) {
//　　for {
//　　　　value, ok := <-ch
//　　　　if ok {
//　　　　　　fmt.Printf("id: %d, recv: %d\n", id, value)
//　　　　} else {
//　　　　　　fmt.Printf("id: %d, closed\n", id)
//　　　　　　break
//　　　　}
//　　}
//　　done <- true
//}
//
//func main() {
//　　ch := make(chan int, 3)
//
//　　coNum := 2
//　　done := make(chan bool, coNum)
//　　for i := 1; i <= coNum; i++ {
//　　　　go Consumer(i, ch, done)
//　　}
//
//　　go Producer(ch)
//
//　　for i := 1; i <= coNum; i++ {
//　　　　<-done
//　　}
//}
