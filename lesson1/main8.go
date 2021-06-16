package main

import "fmt"

func main() {

	nums := []int{2, 3, 4}
	sum := 0

	/**
	Go 语言中 range 关键字用于for循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。
	在数组和切片中它返回元素的索引值，在集合中返回 key-value 对的 key 值。
	*/
	for index, num := range nums {

		sum += num
		fmt.Println("index:", index, "num:", num)
	}

	fmt.Println("sum:", sum)

	for i, num := range nums {

		if num == 3 {

			fmt.Println("index:", i)
		}
	}

	keys := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range keys {

		fmt.Printf("%s->%s\n", k, v)
	}

	value, isExists := keys["c"]
	fmt.Println(value, isExists)

	for i, c := range "go" {

		fmt.Println(i, c)
	}

}
