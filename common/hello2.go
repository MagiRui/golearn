package common

import "fmt"

func sayHellow() {

	fmt.Println("hello2")
	SayHello()
}

func init() {

	fmt.Println("init3")
}
