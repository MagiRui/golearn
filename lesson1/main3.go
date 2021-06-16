package main

import (
	"fmt"
	"math"
)

func max(num1, num2 int) int {

	var result int
	if num1 > num2 {

		result = num1
	} else {

		result = num2
	}

	return result
}

func swap(x, y string) (string, string) {

	return y, x
}

func getSequence() func() int {

	i := 0
	return func() int {

		i += 1
		return i
	}
}

type Circle struct {
	radius float64
}

func (c Circle) getArea() float64 {

	return 3.14 * c.radius * c.radius
}

func getAverage(arr []int, size int) float32 {

	var i, sum int
	var avg float32

	for i = 0; i < size; i++ {

		sum += arr[i]
	}

	avg = float32(sum) / float32(size)
	return avg

}

func main() {

	var a int = 100
	var b int = 200
	var ret int
	ret = max(a, b)
	fmt.Printf("最大值是:%d\n", ret)

	as, bs := swap("Mahesh", "Kumar")
	fmt.Println(as, bs)

	getSquareRoot := func(x float64) float64 {

		return math.Sqrt(x)
	}

	fmt.Println(getSquareRoot(16))

	nextNumber := getSequence()
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())

	var c1 = Circle{

		radius: 10.00,
	}

	fmt.Println("Area of Circle(c1) = ", c1.getArea())

	var balance = []int{1000, 2, 3, 17, 50}
	var avg = getAverage(balance, 5)
	fmt.Printf("平均値为:%f", avg)

}
