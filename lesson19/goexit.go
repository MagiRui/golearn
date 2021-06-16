package main

import (
	"fmt"
	"runtime"
)

func main() {

	//Goexit时defer依然会执行

	exit := make(chan int)
	go func() {

		defer close(exit)
		defer fmt.Println("a")

		func() {

			defer func() {

				fmt.Println("b", recover() == nil)
			}()

			func() {

				fmt.Println("c")
				runtime.Goexit()
				fmt.Printf("c exit")
			}()

			fmt.Println("b exit")

		}()

		fmt.Println("c exit")
	}()

	<-exit
	fmt.Printf("main exit")
}
