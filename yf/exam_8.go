package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "hello"
	boolean := strings.Contains(str, "hell")
	if boolean {
		fmt.Println("yes")
	}
}
