package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan int, 10)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("发送了{}元素", i)
			time.Sleep(500 * time.Millisecond)
			ch1 <- i
		}
		close(ch1)
	}()

	for {
		ele,ok := <- ch1
		if !ok {
			fmt.Println("结束了")
			//break
		}
		time.Sleep(500 * time.Millisecond)
		fmt.Println("接收了{}元素", ele)
	}
	
}
