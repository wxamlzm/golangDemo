package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("1")
	fmt.Println("A")
	defer fmt.Println("2")
	fmt.Println("B")
	defer fmt.Println("3")
	fmt.Println("C")

	foo()
}

func foo() int {
	f, err := os.Open("main.go")
	if err != nil {
		fmt.Println(err)
		return -1
	}

	defer func() {
		fmt.Println("Close file...")
		f.Close()
	}()

	b := make([]byte, 1024)

	n, err := f.Read(b)

	if err != nil {
		fmt.Println(err)
		return -1
	}

	fmt.Println(n)
	fmt.Println(string(b))

	//f.Close()

	return 0

}
