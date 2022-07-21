package main

import (
    "fmt"
)

type LengthConstraint interface {
    string | []byte | []int
}

func length[T LengthConstraint](t T) int {
    return len(t)
}

func main() {
    fmt.Printf("Type constraints\n")
    fmt.Println(length([]int{1,2,3,4,5}))
    fmt.Println(length("Fernando"))
    fmt.Println(length([]byte("Giorgetti")))
}

