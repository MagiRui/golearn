package main

import (
	"fmt"
)

func main() {

	var x, y, z int = 5, 6, 7
	fmt.Println("变量的定义:", x, y, z)

	const (
		a = iota
		b
		c
		d = "ha"
		e
		f = 100
		g
		h = iota
		i
	)

	fmt.Println("常量const及iota", a, b, c, d, e, f, g, h, i)

	var av int = 4
	var ptr *int
	ptr = &av
	fmt.Println("指针", av, ptr, *ptr)

	var grade string = "B"
	var marks int = 90

	switch marks {

	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"

	}

	switch {

	case grade == "A":
		fmt.Println("优秀!")
	case grade == "B", grade == "C":
		fmt.Println("良好!")
	case grade == "D":
		fmt.Println("及格!")
	case grade == "F":
		fmt.Println("不及格")
	default:
		fmt.Println("差")
	}

	fmt.Println("你的等级是:", grade)

	var xType interface{}

	switch i := xType.(type) {

	case nil:

		fmt.Printf("xType类型是：%T", i)
	case int:

		fmt.Printf("xType的类型是int")
	case float64:

		fmt.Printf("xType的类型是float64")
	case func(int) float64:

		fmt.Printf("xType的类型是func(int)")
	case bool, string:

		fmt.Printf("xType是bool或string类型")
	default:

		fmt.Printf("未知类型")

	}

	var C, c1 int
	C = 1
A:
	for C < 100 {

		C++
		for c1 = 2; c1 < C; c1++ {

			if C%c1 == 0 {

				goto A
			}
		}

		fmt.Println(C, "是素数")

	}

}
