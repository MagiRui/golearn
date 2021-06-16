package main

import "fmt"

func main() {

	var a1 [3]int
	var a2 [3]int

	a1 = a2
	fmt.Println(a1)

	b := [4]int{1, 2, 3, 4}
	fmt.Println(b)

	//动态可扩容的数组
	c := [...]int{1, 2, 3}
	c[0] = 5
	fmt.Println(c)

	type stu struct {
		name string
		age  int
	}

	s := stu{

		name: "www",
		age:  11,
	}

	d := [3]stu{

		s,
	}

	fmt.Println(d)

	d2 := [3]stu{

		{name: "xww1", age: 11},
		{name: "xww2", age: 12},
		{name: "xww3", age: 13},
	}

	fmt.Println(d2)

	e := [2][2]int{

		{1, 2}, {3, 4},
	}

	fmt.Println(e)

	//不固定行数的数组
	e1 := [...][2]int{

		{1, 2}, {2, 3}, {3, 4},
	}

	fmt.Println(e1)

	//数组的长度
	fmt.Println(len(d))
	fmt.Println(len(e))
	fmt.Println(len(e1))

	//指针数组  数组指针

	x, y := 10, 20
	x, y = y, x
	z := [...]*int{&x, &y}
	p := &z

	fmt.Println(*p)
	fmt.Println(&x, &y)

	f := [2]int{11, 22}

	var g [2]int
	g = f

	fmt.Printf("%p, %v \n", &f, f)
	fmt.Printf("%p, %v \n", &g, g)

	// g, f, test函数形式参数都不是一份新的拷贝， 如果
	//数组的容量很大，则拷贝效率就会很低，此时可以传递数组
	//的指针
	test(f)
}

func test(x [2]int) {

	fmt.Printf("%p, %v \n", &x, x)
}
