package main

import (
	_ "Init/Math"
	"fmt"
)

func main() {
	a, b := 123, 456
	fmt.Printf("%d+%d = %d\n", a, b, math.Add(a, b))
	fmt.Printf("%d-%d = %d\n", a, b, math.Sub(a, b))
}
