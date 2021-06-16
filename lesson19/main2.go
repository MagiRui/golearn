package main

import (
	"fmt"
	"runtime"
)

func main() {

	runtime.GOMAXPROCS(1)

	exit := make(chan int)
	go func() {

		defer func() {

			close(exit)
		}()

		go func() {

			fmt.Println("B")
		}()

		fmt.Println("A1")
		runtime.Gosched()
		fmt.Println("A2")

	}()

	<-exit

}
