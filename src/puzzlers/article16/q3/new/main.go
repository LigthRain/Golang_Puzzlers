package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var shunxu int64

func main() {
	num := 10

	for i := 0; i < num; i++ {
		go func(i int) {
			fn := func(i int64) {
				fmt.Println(i)
			}
			trigger(int64(i), fn)
		}(i)
	}

	trigger(10, func(i int64) {

	})
}

func trigger(i int64, fn func(i int64)) {
	for  {
		if n := atomic.LoadInt64(&shunxu);n == i {
			fn(i)
			atomic.AddInt64(&shunxu, 1)
			break
		}
		time.Sleep(time.Nanosecond)
	}
}
