package main

import "fmt"

type data struct {
	x int
	s string
}

func main() {

	/*
			1: * / % <<  >> &
			2: + - | ^
		    3: &&
		    4: ||
		    5: ;
	*/

	const (
		read byte = 1 << iota
		write
		exec
		freeze
	)

	/*
		异或，英文为exclusive OR，缩写成xor
		异或（xor）是一个数学运算符。它应用于逻辑运算。异或的数学符号为“⊕”，计算机符号为“xor”。其运算法则为：
		a⊕b = (¬a ∧ b) ∨ (a ∧¬b)
		如果a、b两个值不相同，则异或结果为1。如果a、b两个值相同，异或结果为0。
		异或也叫半加运算，其运算法则相当于不带进位的二进制加法：二进制下用1表示真，0表示假，则异或的运算法则为：0⊕0=0，1⊕0=1，0⊕1=1，1⊕1=0（同为0，异为1），这些法则与加法是相同的，只是不带进位。
	*/
	a := read | write | freeze
	b := read | write | exec
	/*
		 &^将运算符左边数据相异的位保留，相同位清零。
		 1、如果右侧是0，则左侧数保持不变
		 2、如果右侧是1，则左侧数一定清零
		 3、功能同a&(^b)相同
		 4、如果左侧是变量，也等同于：
			var a int
		 	a &^= b
			和它等价的c语言运算符表达式：
			等价于c语言里的&=~
			例如c语言的写法：
				int a = 3;
				a &= ~1;
	*/

	c := a &^ b
	fmt.Printf("%04b &^ %04b = %04b\n", a, b, c)

	x := 10

	var p *int = &x
	*p += 20
	fmt.Println(p, *p)

	/**
	在Go中 ++ 和 -- 只能作为语句而非表达式
	*/
	ax := 1
	ax++
	// c := ax++ 报错，不能是表达式

	var a1 data = data{x: 1, s: "abc"}
	b1 := data{
		1,
		"abc",
	}

	fmt.Println(a1)
	fmt.Println(b1)

	//c1是一个指针
	c1 := &data{

		x: 1,
		s: "abc",
	}

	fmt.Println(b1, c1)

}
