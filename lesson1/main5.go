package main

import "fmt"

const MAX = 3

func swap1(x *int, y *int) {

	var temp int
	temp = *x
	*x = *y
	*y = temp
}

func main() {

	a := []int{10, 100, 200}

	var i int
	var ptr [MAX]*int

	for i = 0; i < MAX; i++ {

		fmt.Printf("a[%d] = %d\n", i, a[i])
	}

	for i = 0; i < MAX; i++ {

		ptr[i] = &a[i]
	}

	for i = 0; i < MAX; i++ {

		fmt.Printf("a[%d] = %d\n", i, *ptr[i])
	}

	var ax int
	var ptrx *int
	var pptrx **int

	ax = 3000
	ptrx = &ax
	pptrx = &ptrx

	fmt.Printf("变量a=%d\n", a)
	fmt.Printf("指针变量 *ptr=%d\n", *ptrx)
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptrx)

	var ax1 int = 100
	var bx1 int = 200

	fmt.Printf("交换前ax的值:%d\n", ax1)
	fmt.Printf("交换前bx的值:%d\n", bx1)

	swap(&ax1, &bx1)
	fmt.Printf("交换前ax的值:%d\n", ax1)
	fmt.Printf("交换前bx的值:%d\n", bx1)

}
