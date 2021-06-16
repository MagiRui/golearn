package main

import (
	"fmt"
	"sync"
	"time"
)

func task1() {

	fmt.Println("task1 start")
	time.Sleep(1 * time.Second)
	fmt.Println("task1 end")
}

func task2() {

	fmt.Println("task2 start")
	time.Sleep(2 * time.Second)
	fmt.Println("task2 end")
}

func main2() {

	// gorouting 协程 或者 线程 goruntime

	//顺序执行
	fmt.Println(time.Now())

	//创建int型的通道
	exit := make(chan int)

	//go加上去之后函数task1就变成一个线程了
	//go task1()
	//go task2()
	go func() {

		fmt.Println("task3 start")
		time.Sleep(1 * time.Second)
		fmt.Println("task3 end")

		exit <- 1
	}()

	//time.Sleep(3*time.Second)
	fmt.Println(time.Now())

	a := <-exit
	fmt.Println(a)

}

func main() {

	// gorouting 协程 或者 线程 goruntime

	//顺序执行
	fmt.Println(time.Now())

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {

		fmt.Println("task3 start")
		time.Sleep(1 * time.Second)
		fmt.Println("task3 end")
		wg.Done()
	}()

	fmt.Println(time.Now())

	//wg 中的数字必须是零的时候才不会阻塞
	wg.Wait()

}
