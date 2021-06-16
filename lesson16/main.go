package main

import "fmt"

//接口只包含方法申明，不能有实现

type tester interface {
	test()
	string() string
}

type data struct {
}

func (d *data) test() {

	fmt.Println("test")
}

func (d *data) string() string {

	fmt.Println("string")
	return "string"
}

func main() {

	d := &data{}
	var i tester
	i = d

	i.test()
	i.string()

	d2 := data{}
	d2.test()
	d2.string()
	/*var i2 tester
	//cannot use d2 (type data) as type tester in assignment:
	//data does not implement tester (string method has pointer receiver)
	i2 = d2
	i2.string()
	i2.test()*/

}
