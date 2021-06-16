package main

import (
	"errors"
	"fmt"
)

func check(x int) error {

	if x <= 0 {

		return errors.New("x<0")
	}

	return nil
}

func main() {

	if x := 3; x >= 3 {

		fmt.Println(x)
	}

	b := 9

	if b > 5 {

		fmt.Println(b)
	} else if b > 7 {

		fmt.Println(b)
	}

	c := -3
	if err := check(c); err == nil {

		c++
		fmt.Println(c)
	} else {

		fmt.Println(err)
	}

	for i := 0; i < 4; i++ {

		fmt.Println(i)
	}

	/*
		for ;;{

			fmt.Println("asdfasf")
		}*/

	//map channel slice  array 支持迭代

	//int a[3] = {1,2,3} in C language
	var data [3]int = [3]int{10, 20, 30}
	//range 后面是可迭代的类型
	for i, s := range data {

		fmt.Println(i, s)
	}

	for i := range data {

		fmt.Println(i, data[i])
	}

	data2 := [3]string{"aa", "bb", "cc"}
	var data2Copy [3]*string
	for i, s := range data2 {

		//fmt.Println(data2[i])
		data2Copy[i] = &s
	}

	for _, s := range data2Copy {

		fmt.Println(*s)
	}

	// goto break continue

	//start:
	for i := 0; i < 4; i++ {

		fmt.Println(i)
		if i > 2 {

			goto end
		}
	}

end:
	fmt.Println("end")

	testSwitch()

}

func test() {

	fmt.Println(" ")
}

func testSwitch() {

	x := 3
	a, b, c := 1, 2, 3

	/*
			default也可以放在第一个case前面，不影响switchy的逻辑
			case b:
			case c:
				fmt.Println("x=c")
		    如上的case语句，如果匹配上了b则什么也不会打印，因为case b:后面是个空语句，且不会继续往下走

			正常的case语句，只会匹配一个case并执行,如果要继续执行后面的一个则添加一个fallthrough即可
	*/
	switch x {

	case a:

		fmt.Println("x=a")
	case b:

		fmt.Println("x=b")
	case c:

		fmt.Println("x=c")
		fallthrough
	default:

		fmt.Println("All not Equal")

	}
}
