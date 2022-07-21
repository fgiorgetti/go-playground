package main

import (
	"fmt"
)

// This playground uses a development build of Go:
// devel go1.18-a174638a5c Fri Dec 3 14:28:11 2021 +0000

func Print[T any](s ...T) {
	for _, v := range s {
		fmt.Print(v)
	}
}

func main() {
	Print("Hello, ", "playground\n")
	Print(10, 20, 30)
	fmt.Println()
	Print(11.1, 22.2, 33.3)
	fmt.Println()
}

