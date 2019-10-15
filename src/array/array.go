package main

import "fmt"

func main() {
	var A [10]int

	for i := 0; i < 10; i++ {
		A[i] = i * i
	}
	fmt.Println(A)

	s := "hello 월드"

	// byte 배열을 사용할 경우 바이트 단위로 순회하면서 프린트
	// rune 배열을 사용할 경우 문자 단위로 순회하면서 프린트
	s2 := []rune(s)
	fmt.Println("len(s2) = ", len(s2))
	for i := 0; i < len(s2); i++ {
		fmt.Print(string(s2[i]), ", ")
	}

	arr := [5]int{1, 2, 3, 4, 5}
	clone := [5]int{}

	for i := 0; i < 5; i++ {
		clone[i] = arr[i]
	}
	fmt.Println("\n", clone)

	arr1 := [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(arr)/2; i++ {
		arr1[i], arr1[len(arr)-1-i] = arr1[len(arr)-1-i], arr[i]
	}
	fmt.Println(arr1)
}
