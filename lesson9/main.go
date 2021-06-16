package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
)

func main() {

	a := func(s string) {

		fmt.Println("main inner func", s)
	}

	a("main inner")

	t3 := test3
	result := test2(t3)
	fmt.Println("func name is param for func", result)

	x, y := 1, 2
	defer func() {

		fmt.Println(x, y)
	}()

	fmt.Println(x, y+1)

}

func test() {

	fmt.Println("test func")
}

//匿名函数
func test2(f func(s string) (int, error)) int {

	a, _ := f("abc")
	return a
}

func test3(s string) (int, error) {

	fmt.Println(s)
	return 2, errors.New("error")
}

func test4() {

	f, err := os.Open("filename")

	//函数执行完之后，执行
	defer f.Close()
	if err != nil {

		fmt.Println(err)
	}

	var c []byte
	b := bytes.NewBuffer(c)
	content, _ := f.Read(b.Bytes())
	fmt.Println(content)
}
