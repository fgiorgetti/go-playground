package main

import (
	"fmt"
)

type Node[T any] struct {
	Value T
}

func main() {
	// initializing generic types
	nodei := Node[int]{
		Value: 17,
	}
	fmt.Printf("%v\n", nodei.Value)

	nodes := Node[string]{
		Value: "Blabla",
	}
	fmt.Printf("%v\n", nodes.Value)
}

