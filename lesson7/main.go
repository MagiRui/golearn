package main

import "fmt"
import "./pkg"

/**
1.函数无需前置申明
2.不支持命名嵌套定义,函数内部可以定义匿名函数
3.具备相同签名的函数视为同一类型
go build -gcflags "-l -m"
*/
func main() {

	fmt.Println("sayHello")
	sayHello()

	fmt.Println("sayHello2")
	sayHello2(2)

	funcSayHello := sayHello
	funcSayHello3 := sayHello3
	funcSayHello = funcSayHello3
	funcSayHello()

	sayHello4ReturnValue := sayHello4()

	fmt.Println("sayHello4ReturnValue:", sayHello4ReturnValue)

	sayHello5ReturnValue := sayHello5()

	fmt.Println("sayHello5ReturnValue:", sayHello5ReturnValue, *sayHello5ReturnValue)

	pkg.Smile()
}

func sayHello() {

	fmt.Println("Hello")

	a := func() {

		fmt.Println("222")
	}

	a()
}

func sayHello2(b int) {

	fmt.Println("Hello")

	a := func() {

		fmt.Println("222")
	}

	a()
}

func sayHello3() {

	fmt.Println("sayHello3 ")

	a := func() {

		fmt.Println("222")
	}

	a()
}

func sayHello4() int {

	b := 9
	return b
}

func sayHello5() *int {

	c := 19
	return &c
}
