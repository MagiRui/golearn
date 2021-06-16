package main

import "fmt"

type stringer interface {
	string() string
}

type tester interface {
	stringer
	test()
}

type data struct {
}

func (d *data) test() {

}

func (d *data) string() string {

	return "aaa"
}

func main2() {

	var d data
	var t tester = &d

	fmt.Println(t.string())
	t.test()

	var d2 data
	var t2 tester = &d2
	pp(t2)

	var s stringer = t
	pp(s)

	/**
	cannot use s (type stringer) as type tester in assignment:
	stringer does not implement tester (missing test method)
	var t3 tester = s
	*/

}

func pp(a stringer) {

	fmt.Println(a.string())
}

func main() {

	//var t int
	//t = 4
	//
	//
	//var s string
	//s = "aaa"

	var i interface{}
	i = 4
	fmt.Println(i)

	i = "abc"
	fmt.Println(i)

	m := make(map[string]interface{}, 0)
	m["key1"] = "value1"
	m["key2"] = "value2"
	fmt.Println(m)

	s := stu{

		name: "xww",
		age:  11,
	}

	m["key3"] = s
	fmt.Println(m)

	d := m["key1"]
	fmt.Println(d)

	value, consits := m["key1"]
	fmt.Println(value, consits)

	/**
	cannot use d (type interface {}) as type int in assignment: need type assertion

	var d1 int
	d1 = d
	*/

	//类型断言
	dInt := m["key1"].(string)
	fmt.Println(dInt)

	var s1 stu
	s1 = m["key3"].(stu)
	fmt.Println(s1)

	// ok 代表转化是成功了还是没成功
	d2, ok := m["key5"].(string)
	fmt.Println(d2, ok)

}

type stu struct {
	name string
	age  int
}
