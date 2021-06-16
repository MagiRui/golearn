package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {

	a := 100
	p := &a

	fmt.Printf("%p  %v\n", &p, p)

	test(p)

	fmt.Println(a)

	test2("aaa", 1, 2, 3, 4, 5, 6)

	var a2 = []int{1, 2, 3, 4}

	//将数组a2转化为切片，再调用test2函数
	test2("aaa", a2...)

	fmt.Println(reflect.TypeOf(a2))
	b := a2[:]
	fmt.Println(reflect.TypeOf(b))

	c, _ := test3(5)
	fmt.Println("test3 return value:", c)
}

//go的函数传值还是传指针？传值
func A() {

}

func test(x *int) {

	*x = 10
	fmt.Printf("%p %v\n", &x, x)
}

func test2(s string, a ...int) { //a is slice

	fmt.Printf("%T %v\n", a, a)

}

func test3(a int) (int, error) {

	b := a + 3
	return b, errors.New("error")
}

func test4(a int) (age int, name string) {

}
