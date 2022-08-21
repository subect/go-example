package main

import "fmt"

func main() {
	s := "pprof.cn博客"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v(%c)\t", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s {
		fmt.Printf("%v(%c)\t", r, r)
	}
	fmt.Println()
	fmt.Println(len(s))
}
