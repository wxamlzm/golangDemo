package main

import "fmt"

func main() {
    type str string
    var a str = "I'm a string"
    fmt.Printf("%T, %s\n", a, a)

    type Student struct {
        name   string
        age    int
        gender string
        score  float32
    }

    b := Student{"张飞", 25, "男", 123.45}
    fmt.Println(b)
    fmt.Println(b.name, b.age, b.gender, b.score)

    c := Student{
        gender: "男",
        name:   "赵云",
    }
    fmt.Println(c)

    var d *Student = &c
    fmt.Println(d, *d, d.name, (*d).gender)

    e := struct {
        x int
        y int
    }{y: 456, x: 123}
    fmt.Println(e)
}
