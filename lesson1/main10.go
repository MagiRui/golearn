package main

import "fmt"

func Factorial(n uint64) (result uint64) {

	if n > 0 {

		result = n * Factorial(n-1)
		return result
	}

	return 1
}

func fibonacci(n int) int {

	if n < 2 {

		return n
	}

	return fibonacci(n-2) + fibonacci(n-1)
}

func main() {

	var i int = 15
	fmt.Printf("%d的阶乘是%d\n", i, Factorial(uint64(i)))

	var ai int
	for ai = 0; ai < 10; ai++ {

		fmt.Printf("%d\t", fibonacci(ai))
	}
}
