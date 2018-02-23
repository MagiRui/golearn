package main

import (
	"fmt"
	"strconv"
)

//变量大写表明可以导入的
//大写的结构体，方法， 可以被别的包引用
type Student struct {
	Name string
	age  int
}

func main() {

	fmt.Printf("Hello World")

	var vs_age int
	fmt.Println(vs_age)

	x, _ := strconv.Atoi("12")

}
