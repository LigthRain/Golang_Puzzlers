package main

import (
	"fmt"
	"time"
)

type FutureJob struct {
	//task
	task func() interface{}
	//通知的chan
	trigger chan interface{}
}

func NewFutureJob(task func() interface{}) *FutureJob {
	return &FutureJob{task: task, trigger: make(chan interface{})}
}

func (f FutureJob) DoTask() {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("error")
			}
		}()
		f.trigger <- f.task()
	}()
}

func (f FutureJob) GetResult() interface{} {
	return <- f.trigger
}

func justAFunction(id int) int {
	return  id * id
}

func main() {
	funct := func(id int) func() interface {} {
		return func() interface{} {
			time.Sleep(time.Second * 5)
			fmt.Println("经历过网络调用5s，最终结果达到了")
			return id * id
		}
	}

	future := NewFutureJob(funct(10))
	future.DoTask()
	fmt.Println("中间任务 耗时耗时")
	fmt.Println(future.GetResult())
	time.Sleep(time.Second * 10)

}
