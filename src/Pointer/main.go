package main

import "fmt"

func main() {
	a := 10
	b := &a

	// 取地址和解引用的语义
	fmt.Println(b, *b)

	*b = 20
	fmt.Println(a)

	var c **int = &b
	fmt.Println(c, *c, **c)

	d := &a
	fmt.Println(b == d, b != d)

	// 动态分配，编译器运行到这里才分配，那相对其他的不同在哪里，其他的一早就分配的么
	e := new(string)
	// e会是个指针
	*e = "beijing"
	fmt.Println(e, *e)

	h := foo()
	fmt.Println(h, *h)

	h = nil
	// 用来判断一个指针是不是为空
	fmt.Println(h, h == nil)
}

// 内存逃逸实验
func foo() *int {
	f := 30
	g := &f

	return g
}
