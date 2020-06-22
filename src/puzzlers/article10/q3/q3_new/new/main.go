package main

func main() {
	ch1 := make(chan int, 1)

	ch1 <- 10
	ch1 <- 10
	//ch1 <- 10
	//fmt.Println( <- ch1)
	//close(ch1)
	//fmt.Println( <- ch1)
}

func SendInt(ch1 chan<- int) {
	ch1 <- 10
	//ele := <-ch1
}
