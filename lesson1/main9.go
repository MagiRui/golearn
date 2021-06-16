package main

import "fmt"

func main() {

	/**
	Map 是一种无序的键值对的集合。Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。
	Map 是一种集合，所以我们可以像迭代数组和切片那样迭代它。不过，Map 是无序的，我们无法决定它的返回顺序，这是因为 Map 是使用 hash 表来实现的。
	*/

	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)

	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"

	for country := range countryCapitalMap {

		fmt.Println("Capital of ", country, "is", countryCapitalMap[country])
	}

	/* 查看元素在集合中是否存在 */
	capital, ok := countryCapitalMap["United states"]
	/* 如果 ok 是 true, 则存在，否则不存在 */

	if ok {

		fmt.Println("Capital of United States is", capital)
	} else {

		fmt.Println("Capital of United States is not present")
	}

	/*
		delete() 函数
		delete() 函数用于删除集合的元素, 参数为 map 和其对应的 key
	*/

	delete(countryCapitalMap, "France")
	fmt.Println("Entry for France is deleted")

	fmt.Println("删除元素后map")

	for country := range countryCapitalMap {

		fmt.Println("Capital of ", country, "is", countryCapitalMap[country])
	}

	fmt.Println("key value for range:")
	for country, capital := range countryCapitalMap {

		fmt.Println(country, capital)
	}

}
