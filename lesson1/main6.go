package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {

	var Book1 Books
	var Book2 Books

	Book1.title = "Go Language"
	Book1.author = "www.3apple.cn"
	Book1.subject = "Go Learning"
	Book1.book_id = 20180308180659

	Book2.title = "Python Language"
	Book2.author = "www.3apples.cn"
	Book2.subject = "Python Learning"
	Book2.book_id = 20180308163102

	fmt.Println(Book1)
	fmt.Println(Book2)

}
