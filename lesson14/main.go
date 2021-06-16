package main

import (
	"fmt"
	"unsafe"
)

type node struct {
	id  int32
	age int8
	tel int16
}

type user struct {
	name string
	age  int
}

type attr struct {
	owner int
	perm  int
}

type file struct {
	name string
	//
	attr
}

func main() {

	//fmt.Println(unsafe.Sizeof(node))

	u := user{

		name: "MagiRui",
		age:  11,
	}

	fmt.Println(u)
	u2 := user{"MagiRui1", 12}
	fmt.Println(u2)

	u3 := struct {
		name string
		age  int
	}{
		name: "Go Language",
		age:  8,
	}

	fmt.Println(u3)

	/*type file struct {
		name string
		attr struct {
			owner int
			perm  int
		}
	}

	f := file{

		name: "aa.txt",
	}

	f.attr.owner = 1
	f.attr.perm = 2*/

	f2 := file{

		name: "aa.txt",
		attr: attr{

			owner: 1,
			perm:  2,
		},
	}

	fmt.Println(f2, f2.perm)

	//空结构体
	var a struct{}
	var b [100]struct{}
	fmt.Println(unsafe.Sizeof(a), unsafe.Sizeof(b), len(b))

}
