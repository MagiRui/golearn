package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"time"
)

func count() {

	x := 0
	for i := 0; i < math.MaxInt32; i++ {

		x += 1
	}

	fmt.Println(x)
}

func test1(n int) {

	for i := 0; i < n; i++ {

		count()
	}
}

func test2(n int) {

	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {

		go func() {

			count()
			wg.Done()
		}()
	}
	wg.Wait()
}

func main() {

	//GOMAXPROCS

	runtime.GOMAXPROCS(runtime.NumCPU())

	fmt.Println(runtime.NumCPU())

	//有多少个Cpu,就用多少个Cpu
	n := runtime.GOMAXPROCS(0)
	start := time.Now()
	test2(n)
	end := time.Now()
	fmt.Println(start, end)
}
