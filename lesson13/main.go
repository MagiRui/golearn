package main

import (
	"fmt"
	"time"
)

func main() {

	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	m["c"] = 3
	m["d"] = 4
	m["e"] = 5

	fmt.Println(m)

	m["a"] = 3
	fmt.Println(m)

	v, ok := m["f"]

	fmt.Println(v, ok)
	if ok {

		fmt.Println(v, ok)
	}

	m["abc"] = 123

	fmt.Println(len(m))

	for k, v := range m {

		fmt.Println(k, v)
	}

	go func() {

		for {

			m["a"] = 1
			time.Sleep(time.Microsecond)
		}
	}()

	go func() {

		for {

			_ = m["a"]
			time.Sleep(time.Microsecond)
		}
	}()

	select {}

}
