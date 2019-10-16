package main

import "fmt"

type Student struct {
	name string
	age  int

	grade string
	class string
}

func (s Student) PrintSungjuk() {
	fmt.Println(s.class, s.grade)
}

func (s *Student) InputSungjuck(class string, grade string) {
	s.class = class
	s.grade = grade
}

func main() {

	s := Student{name: "철수", age: 30, class: "수학", grade: "A"}
	s.InputSungjuck("과학", "C")
	s.PrintSungjuk()

	var p *int
	var a int

	p = &a
	a = 3

	fmt.Println(a)
	fmt.Println(p)
	fmt.Println(*p)

	var b int

	b = 2
	p = &b

	fmt.Println(b)
	fmt.Println(p)
	fmt.Println(*p)

	Increase(p)

	fmt.Println(b)

}

func Increase(x *int) {
	*x = *x + 1
}
