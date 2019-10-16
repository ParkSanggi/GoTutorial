package main

import "fmt"

type Student struct {
	name  string
	class int

	achievement Achievement
}

type Achievement struct {
	name  string
	grade string
}

func (s Student) ViewAchievement() {
	fmt.Println(s.achievement)
}

func main() {
	var s Student
	s.name = "철수"
	s.class = 1

	s.achievement.name = "수학"
	s.achievement.grade = "c"

	s.ViewAchievement()
}
