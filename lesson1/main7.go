package main

import (
	"fmt"
)

func printSlice(x []int) {

	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func main() {

	//make([]T, length, capacity)
	var nubers = make([]int, 3, 5)
	printSlice(nubers)

	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for x := range numbers {

		fmt.Println("for range slice", x)
	}

	printSlice(numbers)
	number1 := make([]int, 0, 5)
	printSlice(number1)

	number2 := numbers[:2]
	printSlice(number2)

}

func main3() {

	//append() 和 copy() 函数
	//如果想增加切片的容量，我们必须创建一个新的更大的切片并把原分片的内容都拷贝过来。
	//下面的代码描述了从拷贝切片的 copy 方法和向切片追加新元素的 append 方法。

	var numbers []int
	printSlice(numbers)

	numbers = append(numbers, 0)
	printSlice(numbers)

}
