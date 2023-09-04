package main

import "fmt"

func main() {
	a := "12345678"
	fmt.Println(a)
	b := `12345678`
	fmt.Println(b)

	c := "123\n456\n789"
	fmt.Println(c)

	d := `123\n456\n789`
	fmt.Println(d)

	e := `123`

	fmt.Println(e)

	f := "123\"456\"780"
	fmt.Println(f)

	g := "123`456`789"
	fmt.Println(g)

	h := `123"456"789`
	fmt.Println(h)

	fmt.Println(len(c), len(d), len(e))

	j := "123四五六789"
	fmt.Println(j, len(j))

	for i := 0; i < len(j); i++ {
		fmt.Printf("%x", j[i])
	}
}
