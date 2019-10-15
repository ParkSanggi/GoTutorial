package main

import "fmt"

func main() {
	for dan := 2; dan <= 9; dan++ {
		if dan == 5 {
			continue
		}
		fmt.Printf("%d단\n", dan)
		for i := 1; i <= 9; i++ {
			if dan == 3 && i == 2 {
				continue
			}
			fmt.Printf("%d X %d = %d\n", dan, i, dan*i)
		}
		fmt.Println()
	}
}
