package main

import (
	_ "./in"
	"fmt"
)

const (
	n1 = iota + 1
	n2
	n3
	n4
	n5
	n6
)

func main() {
	fmt.Println(n1, n2, n3, n4, n5, n6)
}
