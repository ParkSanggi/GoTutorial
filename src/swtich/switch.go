package main

import "fmt"

func main() {
	x := 33

	switch {
	// x를 적지 않을 경우 기본 값 true
	case x > 40:
		fmt.Println("x = 31")
	case x < 20:
		fmt.Println("x = 32")
	case x > 30:
		fmt.Println("x = 33")
	// 위에서부터 순차적으로 확인하므로 더이상 아래로 가지 않고 출력됨
	case x == 33:
		fmt.Println("x는 33")
	}
}
