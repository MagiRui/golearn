package main

import "fmt"

func main() {

	//定义一个10个元素的数组
	//var balance [10] float32

	var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	for xvalue := range balance {

		fmt.Println("for balance:", xvalue)
	}
	fmt.Println(balance)

	var balance2 = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Println(balance2, len(balance2), cap(balance2))

	towDimensionArray := [3][4]int{

		{0, 1, 2, 3},
		{5, 6, 7, 8},
		{4, 9, 10, 11},
	}

	for i := 0; i < 3; i++ {

		for j := 0; j < 4; j++ {

			fmt.Printf("a[%d][%d]= %d\n", i, j, towDimensionArray[i][j])
		}
	}

	var n [10]int
	var i, j int
	for i = 0; i < 10; i++ {

		n[i] = i + 100
	}

	for j = 0; j < 10; j++ {

		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}

	var a int = 10
	var ip *int
	ip = &a
	fmt.Printf("变量的地址:%x\n", &a)

	fmt.Printf("变量存储的指针地址:%x\n", ip)

	fmt.Printf("*ip变量的值:%d\n", *ip)

}
