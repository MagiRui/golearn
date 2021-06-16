package main

import "fmt"

func main(){

	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)

	var xy uint8 = 0
	fmt.Println(xy-1)

	medals := []string{"gold", "silver", "bronze"}

	for i:= len(medals)-1; i>=0; i--{


		fmt.Println(i, medals[i])
	}

	s1 := "Hello, 世界"
	fmt.Println(len(s1))
	for x := range(s1){

		fmt.Println(x)
	}

}


