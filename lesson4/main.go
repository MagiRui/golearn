package main

import (
	"fmt"
	"math"
)

type Student struct {
	Age  int
	Name string
}

func main() {

	/*
			bool
			byte
		int 8 16 32 64
		uint 无符号整型
		float float8 float16 float32 float 64 float128
		complex64 128  1+2i
		uintptr
		string
		array
		struct
		function
		interface

		map
		slice
		channel
	*/

	s := "abc"
	fmt.Println([]byte(s))

	stu1 := &Student{}

	stu2 := new(Student)

	//取地址
	fmt.Println(stu1)
	fmt.Println(stu2)

	//取内容
	fmt.Println(*stu1)
	fmt.Println(*stu2)

	fmt.Println(math.MinInt8, math.MaxInt16, math.MaxUint16)

	//创建Map
	m := make(map[int]string)
	m[1] = "1"
	m[2] = "2"
	fmt.Println(m)

	//创建slice
	s1 := []int{1, 2, 3, 4, 5}
	s1 = s1[2:4]
	fmt.Println(s1)

	s2 := make([]int, 5)
	s2 = append(s2, 12)
	fmt.Println(s2)

	s3 := make([]int, 0)
	s3 = append(s3, 12)
	fmt.Println(s3)

	s4 := make([]int, 1, 5)
	s4 = append(s4, 13)
	s4 = append(s4, 13)
	s4 = append(s4, 13)
	s4 = append(s4, 13)
	s4 = append(s4, 13)
	s4[5] = 19
	fmt.Println(s4)

	//管道
	c := make(chan int, 2)
	c <- 1
	fmt.Println(<-c)

	/*	p:= new(map[int]string)
		m1 := *p
		m1[1]= "Stu"
		fmt.Println(p)
		fmt.Println(m1)*/

	a := 10
	b := uint8(a)
	fmt.Println(b)

	//type color int

	var red int
	red = 1
	var green int
	green = 2

	fmt.Println(red, green)

	const (
		red_color = iota
	)

	//自定义类型
	type color int
	var redColor color
	redColor = 1
	var yellowColor color
	yellowColor = 2

	fmt.Println(redColor, yellowColor)

	//redColor1 := 1
	//虽然底层类型一样，但是却不能直接赋值
	//redColor1  = redColor

}
