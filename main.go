package main

import "fmt"

func main() {
	var count = 1

	if count == 0 {
		fmt.Println("Hello World")
	} else {
		fmt.Println("Goodbye World")
	}

	sum := 0
	for i := 1; i < 5; i++ {
		sum += i
	}
	fmt.Println(sum) // 10 (1+2+3+4)
}
