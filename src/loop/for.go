package main

import "fmt"

func main() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	fmt.Println("최종 i 값", i)

	/* 아래와 같이 전처리 문을 통해서 i를 선언할 수 있다.
	   아래와 같이 선언 할 경우 i의 스코프가 for문에 한정되어
	   for 문 밖에서는 i를 찾을 수 없다.
	*/
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 4까지 출력되고 멈춤
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	// 5와 7이 출력되지 않고 멈춤
	i = 0
	for {
		if i == 5 {
			i++
			// for문을 한 번 건너 뜀
			continue
		}
		if i == 7 {
			break
		}
		fmt.Println(i)
		i++
	}

}
