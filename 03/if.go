package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		// fmt.Println(i)
		if i%2 == 0 {
			// even
			fmt.Println(i, "even")
		} else {
			// odd
			fmt.Println(i, "odd")
		}
	}
}
