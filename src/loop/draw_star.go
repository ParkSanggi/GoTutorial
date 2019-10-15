package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5-i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i := 0; i < 5; i++ {

		switch {
		case i < 3:
			for j := 0; j < i+1; j++ {
				fmt.Print("*")
			}
		case i >= 3:
			for j := 0; j < 5-i; j++ {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
}
