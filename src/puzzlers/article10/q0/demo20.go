package main

func main() {
	//无缓冲的通道，会在一开始发送or接收的情况下就阻塞，直到有对应的操作执行
	ch1 := make(chan int)
	ch1 <- 2
	//ch1 <- 1
	//ch1 <- 3
	//elem1 := <-ch1
	//fmt.Printf("The first element received from channel ch1: %v\n",
	//	elem1)
}
