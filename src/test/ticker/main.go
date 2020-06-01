package main

import (
	"fmt"
	"time"
)

func main() {
	//场景1：
	for {
		select {
		case <- time.After(10 * time.Microsecond):
			fmt.Println("hello timer")
		}
	}

	//场景2：
	for {
		select {
		case <- time.Tick(10 * time.Microsecond):
			fmt.Println("hello, tick")
		}
	}
}
