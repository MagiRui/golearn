package main

import (
	"fmt"
	"sync"
	"time"
)

func main1() {

	//channel

	event := make(chan struct{})
	c := make(chan string)

	go func() {

		s := <-c
		fmt.Println(s)
		close(event)
	}()

	time.Sleep(2 * time.Second)
	c <- "hi"
	d := <-event
	fmt.Println(d)

	ccn := make(chan int, 3)
	ccn <- 1
	ccn <- 2
	fmt.Println(<-ccn)
	ccn <- 3
	ccn <- 4
	fmt.Println(<-ccn)

	c1 := make(chan string, 1)
	c1 <- "abc"

	a1 := make(chan int)
	b1 := make(chan int, 3)
	b1 <- 1
	b1 <- 2

	fmt.Println("a1", len(a1), cap(a1))
	fmt.Println("b1", len(b1), cap(b1))

	//ok 是不是正常取回
	valueB1, ok := <-b1
	fmt.Println(valueB1, ok)
}

func main2() {

	done := make(chan struct{})
	c := make(chan int)

	go func() {

		defer close(done)
		for {

			x, ok := <-c
			if !ok {

				return
			}

			fmt.Println(x)
		}
	}()

	c <- 1
	c <- 2
	c <- 3

	close(c)
	xdone := <-done
	fmt.Println(xdone)

}

func main3() {

	done := make(chan struct{})
	c := make(chan int)
	go func() {

		defer close(done)
		for x := range c {

			fmt.Println(x)
		}
	}()

	c <- 1
	c <- 2
	c <- 3

	close(c)

	<-done
}

func main4() {

	a := make(chan int, 3)

	a <- 1
	a <- 2

	close(a)

	value1, ok1 := <-a
	fmt.Println(value1, ok1)
	value2, ok2 := <-a
	fmt.Println(value2, ok2)
	value3, ok3 := <-a
	fmt.Println(value3, ok3)

}

func main5() {

	var wg sync.WaitGroup
	wg.Add(2)

	c := make(chan int)
	var send chan<- int = c
	var recv <-chan int = c

	go func() {

		defer wg.Done()

		for x := range recv {

			fmt.Println(x)
		}
	}()

	go func() {

		defer wg.Done()
		defer close(c)
		for i := 0; i < 3; i++ {

			send <- i
		}
	}()

	wg.Wait()

	/*d := make(chan int , 2)
	var send2 chan <- int = d
	var recv2 <- chan int = d*/

}

func main6() {

	//select
	var wg sync.WaitGroup
	a, b := make(chan int), make(chan int)
	wg.Add(2)
	go func() {

		defer wg.Done()
		for {
			var name string
			var x int
			var ok bool
			select {

			case x, ok = <-a:
				name = "a"
			case x, ok = <-b:
				name = "b"

			}

			if !ok {

				return
			}

			fmt.Println(name, x)
		}
	}()

	go func() {

		defer wg.Done()
		defer close(a)
		defer close(b)

		for i := 0; i < 10; i++ {

			select {

			case a <- i:
			case b <- i * 10:
			}
		}
	}()

	wg.Wait()

}

func main() {

	t := time.NewTicker(1 * time.Second)
	for {

		select {

		case <-t.C:
			fmt.Println("aa")
		}
	}
}
