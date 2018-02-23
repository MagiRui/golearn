package main

import "fmt"

const x, y int = 1, 0x23
const s = "Hello World"
const c = "aa"
const xc, yc = "123", 345

func main() {

	const d = 123
	fmt.Println("内部d:")
	fmt.Println(d)

	const y = 123
	{
		const x = "aavc"
		fmt.Printf("内部x:")
		fmt.Println(x)
	}

	const (
		x1, y1 = 1, 2
		n      = uint8(12)
		b      = byte(x)
	)

	j := 2
	fmt.Println(&j)

	const k = 2
	//fmt.Println(&k)

	b1 := k + 2
	fmt.Println(&b1)

}
