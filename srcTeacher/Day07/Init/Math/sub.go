package math

import "fmt"

func Sub(x, y int) int {
    return x - y
}

func init() {
    fmt.Println("init-1 in Math/sub.go")
}

func init() {
    fmt.Println("init-2 in Math/sub.go")
}
