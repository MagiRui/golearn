package lesson23

import (
	"fmt"
	"reflect"
	"unsafe"
)

type X int
type Y int

func main1() {

	/*var a X = 100
	t := reflect.ValueOf(a)
	fmt.Println(t.Kind())*/

	var a, b X = 100, 200
	var c Y = 300
	ta, tb, tc := reflect.TypeOf(a), reflect.TypeOf(b), reflect.TypeOf(c)
	fmt.Println(ta == tb, tb == tc)
	fmt.Println(tb.Kind() == tc.Kind())
	fmt.Println(ta)
	fmt.Println(tb)
	fmt.Println(tc)
}

type user struct {
	name string
	age  int
}

type manager struct {
	user
	title string
}

func main2() {

	var m manager
	t := reflect.TypeOf(m)
	if t.Kind() == reflect.Ptr {

		//Elem只用在接口 或 指针类型上，获取所支持的类型
		t = t.Elem()
	}

	for i := 0; i < t.NumField(); i++ {

		f := t.Field(i)
		fmt.Println(f.Name, f.Type, f.Offset)
		if f.Anonymous {

			for x := 0; x < f.Type.NumField(); x++ {

				af := f.Type.Field(x)
				fmt.Println("", af.Name, af.Type, af.Offset)
			}
		}
	}

}

type A int
type B struct {
	A
}

func (A) av() {

}

func (*A) ap() {

}

func (B) bv() {

}

func (*B) bp() {

}

func main3() {

	var b B
	t := reflect.TypeOf(&b)
	s := []reflect.Type{t, t.Elem()}
	for _, t := range s {

		fmt.Println(t, ":")
		for i := 0; i < t.NumMethod(); i++ {

			fmt.Println(" ", t.Method(i))

		}
	}

}

type user2 struct {

	//冒号后面不能有空格，有空格则获取不到
	name int `field:"name" type:"varchar(50)"`
	age  int `field:"age"  type:"int"`
}

func main4() {

	/*var s http.Server
	t := reflect.TypeOf(s)
	fmt.Println("t" , t )
	for i:=0; i< t.NumField();i++{

		fmt.Println("", t.Field(i).Name)
	}*/

	var u user2
	t1 := reflect.ValueOf(u).Type()
	fmt.Println(t1)
	fmt.Println(reflect.TypeOf(u))

	t2 := reflect.TypeOf(u)
	for i := 0; i < t2.NumField(); i++ {

		f := t2.Field(i)
		fmt.Printf("%s:%s %s\n", f.Name, f.Tag.Get("field"), f.Tag.Get("type"))
	}

}

type user3 struct {
	Name string
	code int
}

func main5() {

	a := 100

	va, vp := reflect.ValueOf(a), reflect.ValueOf(&a).Elem()

	fmt.Println(va.CanAddr(), va.CanSet())
	fmt.Println(vp.CanAddr(), vp.CanSet())
	fmt.Println(va.Type(), vp.Type())

	p := new(user3)
	v := reflect.ValueOf(p).Elem()

	name := v.FieldByName("Name")
	code := v.FieldByName("code")

	fmt.Println(name.CanAddr(), name.CanSet())
	fmt.Println(code.CanAddr(), code.CanSet())

	name.SetString("tom")
	fmt.Println(p.Name)

	//code.SetInt(1)
	//fmt.Println(p.code)

	*(*int)(unsafe.Pointer(code.UnsafeAddr())) = 11

	fmt.Println(p.code)
}
