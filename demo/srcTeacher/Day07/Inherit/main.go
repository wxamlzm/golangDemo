package main

import "fmt"

type Human struct {
    name   string
    age    int
    gender string
}

func (h Human) Who() {
    fmt.Printf("我叫%s，今天%d岁，我是%s生。\n",
        h.name, h.age, h.gender)
}

func (h Human) Eat(food string) {
    fmt.Printf("%s吃%s。\n", h.name, food)
}

type Student struct {
    Human
    score float32
}

func (s Student) Who() {
    fmt.Printf("我叫%s，今天%d岁，我是%s生，考了%g分。\n",
        s.name, s.age, s.gender, s.score)
}

func (s Student) Learn(course string) {
    fmt.Printf("%s学%s。\n", s.name, course)
}

type Teacher struct {
    Human
    salary float32
}

func (t Teacher) Who() {
    fmt.Printf("我叫%s，今天%d岁，我是%s生，月薪%g元。\n",
        t.name, t.age, t.gender, t.salary)
}

func (t Teacher) Teach(course string) {
    fmt.Printf("%s教%s。\n", t.name, course)
}

func main() {
    s := Student{
        score: 90,
    }
    s.name = "张飞"
    s.age = 25
    s.gender = "男"
    s.Who()
    s.Eat("包子")
    s.Learn("Go")

    t := Teacher{
        salary: 20000,
    }
    t.name = "赵云"
    t.age = 40
    t.gender = "男"
    t.Who()
    t.Eat("面条")
    t.Teach("C++")

    t.Human.Who()
}
