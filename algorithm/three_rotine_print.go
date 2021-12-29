package algorithm

import "fmt"

func PrintNew() {
	chanNum := 3                           // chan 数量
	chanQueue := make([]chan int, chanNum) // 创建chan Slice
	var result = 0                         // 值
	exitChan := make(chan bool)            // 退出标识
	for i := 0; i < chanNum; i++ {
		//	创建chan
		chanQueue[i] = make(chan int)
		if i == chanNum-1 {
			//	给最后一个chan写一条数据，为了第一次输出从第1个chan输出
			go func(i int) {
				chanQueue[i] <- 1
			}(i)
		}
	}

	for i := 0; i < chanNum; i++ {
		var lastChan chan int //    上一个goroutine 结束才能输出 控制输出顺序
		var curChan chan int  //	当前阻塞输出的goroutine
		if i == 0 {
			lastChan = chanQueue[chanNum-1]
		} else {
			lastChan = chanQueue[i-1]
		}
		curChan = chanQueue[i]
		go func(i int, lastChan, curChan chan int) {
			for {
				if result > 100 {
					//	超过100就退出
					exitChan <- true
				}
				//	一直阻塞到上一个输出完，控制顺序
				<-lastChan
				fmt.Printf("thread%d: %d \n", i, result)
				result = result + 1
				//	当前goroutine已输出
				curChan <- 1
			}
		}(i, lastChan, curChan)

	}
	<-exitChan
	fmt.Println("done")
}
