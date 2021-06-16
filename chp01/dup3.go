package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main(){

	counts := make(map[string]int)
	fmt.Println("dup2...", os.Args[:])
	for _, filename := range os.Args[1:]{

		data, err := ioutil.ReadFile(filename)
		if err != nil {

			fmt.Fprintln(os.Stderr, "dup3:%v\n", err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n"){

			counts[line]++
		}

	}

	for line, n := range counts{

		if n > 1 {

			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
