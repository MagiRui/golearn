package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
终端情况下请使用 ctrl+d
文件的是直接
cat input | go run main.go
 */

func main(){

	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {

		fmt.Println(input.Text())
		counts[input.Text()]++
	}

	for line, n := range(counts){

		if n > 1 {

			fmt.Printf("%d\t%s\n", n , line )
		}

		fmt.Println(counts)

	}
}
