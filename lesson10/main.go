package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	s := "abc\x62\u0041"

	fmt.Println(s)

	var s2 string

	//类型转换错误
	//fmt.Println(s2==nil)
	fmt.Println(s2 == "")
	//转译字符
	//s3 := "abc\\ndef"

	//原生字符
	s3 := `abc\ndef
      def`
	fmt.Println(s3)

	s4 := "abc" + "cde"
	fmt.Println(s4 == "abccde")
	fmt.Println(s4 > "bbc")

	s5 := "abcdefg"
	fmt.Println(s5[0])
	fmt.Printf("%c\n", s5[1])

	//s5[1] ="6" 字符串无法改变

	s6 := s5[1:3]
	fmt.Println(s6)

	s7 := s5[:4]
	fmt.Println(s7)

	fmt.Printf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer(&s5)))
	fmt.Printf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer(&s6)))
	fmt.Printf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer(&s7)))

	//字符串遍历
	for i := 0; i < len(s5); i++ {

		fmt.Printf("%c\n", s5[i])
	}

	for i, c := range s5 {

		fmt.Printf("%d, %c\n", i, c)
	}

	//字符串字节转换
	b := []byte(s5)
	fmt.Println(b)
	b[0] = 98
	s8 := string(b)
	fmt.Println(s8)

	//s9 := make([]byte, 100)

}
