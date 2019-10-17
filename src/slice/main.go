package main

import "fmt"

func main() {
	var s []int

	// cap과 len을 모두 3으로  지정
	s = make([]int, 3)

	s[0] = 100
	s[1] = 200
	s[2] = 300

	/*  해당 slice에 append를 시도할 때 cap이 모자르면
	새 슬라이스를 더 크게만들고 기존 슬라이스의 값을 복사한 뒤 append 한다. */

	s = append(s, 400, 500, 600, 700)
	fmt.Println(s, len(s), cap(s))

	// cap이 모자르지 않으면 append 후 len을 확장
	s = append(s, 400)
	fmt.Println(s, len(s), cap(s))

	arr := []int{1, 2, 3}
	tmp := make([]int, len(arr))
	copy(tmp, arr)

	var p *[]int
	var pp *[]int

	p = &tmp
	pp = &arr

	fmt.Println(&p)
	fmt.Println(&pp)

	fmt.Println(len(tmp), cap(tmp), len(arr), cap(arr))
}
