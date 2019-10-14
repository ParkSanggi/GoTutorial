package main

import "fmt"

func main() {
	a := 4
	b := 2

	fmt.Printf("a&b = %v\n", a&b)
	fmt.Printf("result = %v\n", a|b)
	fmt.Println("result2 = ", a^b)

	var c bool
	c = 3 > 4

	var d bool
	d = 3 < 4 && 2 < 5

	e := 3 < 4 || 2 > 5

	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	if 3 > 5 {
		fmt.Println("참")
	} else if c == true {
		fmt.Println("c")
	} else {
		fmt.Println("거짓")
	}

}
