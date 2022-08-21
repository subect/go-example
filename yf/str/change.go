package main

import "fmt"

func main() {
	s1 := "hello"
	s2 := "博客"
	bytes1 := []byte(s1)
	bytes1[0] = 'H'
	bytes2 := []rune(s2)
	bytes2[0] = '狗'
	fmt.Println(string(bytes1))
	fmt.Println(string(bytes2))
}
