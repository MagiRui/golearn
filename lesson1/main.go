package main

import (
	"fmt"
	"reflect"
)

var x1 = 100

func main() {

	fmt.Printf("Hello World! God Bless You!\n")

	var x int
	var y int
	var z string

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	//自动推断类型
	var a, s = 1000, "asd"
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(s))

	r := 1
	fmt.Println(r)

	xx := 3

	fmt.Println(xx)

	fmt.Println("变量交换")
	x2, y2 := 1, 2
	x2, y2 = y2, x2
	fmt.Println(x2)
	fmt.Println(y2)

}
