package main

import "fmt"

type Shape interface {
	Draw()
}

type Rectangle struct {
	X, Y, W, H int
}

// 由此构成多态
func (r Rectangle) Draw() {
	fmt.Printf("在(%d, %d) 出画一个宽%d高%d的矩形。\n", r.X, r.Y, r.W, r.H)
}
