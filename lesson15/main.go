/*
package main

import "fmt"

//方法是属于类的，函数是独立可以调用的
//方法有关联的状态

type Stu struct {

	name string
	age int
}

type N int

func main(){

	fmt.Println(aa(1))

	var s Stu
	fmt.Println(s.setAge(11))

	var a N = 23
	var b int = 0
	fmt.Println(b)
	fmt.Println(a.toString())

}

func (s *Stu) setAge(a int) int{

	s.age = a
	return s.age
}

func (n N) toString() string{

	return fmt.Sprint("%s", n)
}

func aa(a int) string {

	return "aa"
}
*/

package main

import (
	"fmt"
	"sync"
)

type N int

func (n N) value() {

	n++
	fmt.Printf("v:%p, %v\n", &n, n)
}

func (n *N) Pointer() {

	(*n)++
	fmt.Printf("v:%p, %v\n", &n, *n)
}

type Stu struct {
	age int
}

func (s Stu) value() {

	s.age = 11
}

func (s *Stu) pointer() {

	s.age = 11
}

type people struct {
	name string
}

type Stu2 struct {
	age    int
	people //匿名结构体,此时Stu2中可以直接使用people的属性和方法
}

func (p *people) getname() string {

	return p.name
}

type data struct {
	sync.Mutex
	buf [2014]byte
}

func main() {

	var a N = 23

	a.Pointer()
	a.value()

	fmt.Printf("v:%p, %v\n", &a, a)

	a1 := N(11)

	fmt.Println(a1)

	var s1 Stu

	s1.age = 22
	s1.value()
	fmt.Println(s1.age)

	s1.pointer()
	fmt.Println(s1.age)

	var s2 *Stu
	s2 = &Stu{age: 11}
	s2.pointer()

	fmt.Println(s2.age)

	var s3 Stu2
	fmt.Println("aaa", s3.getname(), "bbb")

	d := data{}
	d.Lock()

	//业务处理

	d.Unlock()

}
