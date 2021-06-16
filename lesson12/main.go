package main

import "fmt"

func main() {

	//slice 是一种数组，和数组是强绑定的
	//Go 语言切片是对数组的抽象。
	//Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go中提供了一种灵活，功能强悍的内置类型切片("动态数组"),与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。

	x := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(len(x))
	fmt.Println(cap(x))

	s := x[2:5]
	fmt.Println(len(s))
	fmt.Println(cap(s))

	fmt.Println(s[0])

	//长度位3,容量为5
	s1 := make([]int, 3, 5)
	fmt.Println(s1)

	var a []int
	b := []int{}
	fmt.Println(a == nil, b == nil)

	fmt.Println(s1)
	s2 := append(s1, 10)
	s2 = append(s2, 10)
	s2 = append(s2, 10)
	fmt.Println(s1)
	fmt.Println(s2)

	s3 := x[2:4]
	fmt.Println(x)
	n := copy(x[3:], s3)
	fmt.Println(n)
	fmt.Println(x)

}
