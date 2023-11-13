package math

import "fmt"

func Add(x, y int) int {
    //init()
    return x + y
}

func init() {
    fmt.Println("init-1 in Math/add.go")
}

func init() {
    fmt.Println("init-2 in Math/add.go")
}
