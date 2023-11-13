package main

import "fmt"

type Student struct {
    name   string
    age    int
    gender string
    score  float32
}

func (s Student) Who() {
    fmt.Printf("我叫%s，今天%d岁，我是%s生，考了%g分。\n",
        s.name, s.age, s.gender, s.score)
    s.score += 0.5
}

func (s *Student) Eat(food string) {
    fmt.Printf("%s吃%s。\n", s.name, food)
}

func (this *Student) Learn(course string) {
    fmt.Printf("%s学%s。\n", this.name, course)
    this.score += 0.5
}

func main() {
    s := Student{
        name:   "张飞",
        age:    25,
        gender: "男",
        score:  90,
    }
    s.Who()
    fmt.Println(s)
    s.Eat("包子")
    s.Learn("Go")
    fmt.Println(s)
}
