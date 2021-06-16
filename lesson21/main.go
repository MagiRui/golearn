package main

import (
	"fmt"
	"sync"
	"time"
)

func main1() {

	//mutex

	var m sync.Mutex
	a := 0

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {

		for i := 0; i < 50000; i++ {

			m.Lock()
			a = a + 1
			m.Unlock()
		}

		wg.Done()
	}()

	go func() {

		for i := 0; i < 50000; i++ {
			m.Lock()
			a = a + 1
			m.Unlock()
		}
		wg.Done()

	}()

	wg.Wait()
	fmt.Println(a)

}

type data struct {
	sync.Mutex
	a int
}

func (d *data) test(s string) {

	d.Lock()
	defer d.Unlock()
	for i := 0; i < 5; i++ {

		fmt.Println(s, i)
		time.Sleep(time.Second)
	}

}

func main() {

	var wg sync.WaitGroup
	wg.Add(2)
	var d data
	go func() {

		defer wg.Done()
		d.test("read")
	}()

	go func() {

		defer wg.Done()
		d.test("write")
	}()

	wg.Wait()
}
