package main

import "fmt"

func main() {

	num := 10
	sing := make(chan struct{}, num)

	for i := 0; i < num; i++ {
		go func() {
			fmt.Println("aaaa")
			sing <- struct{}{}
		}()
	}

	for i := 0; i<num; i++ {
		<- sing
	}

}
