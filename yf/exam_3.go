package main

import "fmt"

func main() {
	str := "hello"
	strByte := []byte(str)
	strByte[0] = 'x'
	strss := string(strByte)
	fmt.Println(strss)
	for _, value := range str {
		if value == 'h' {
			value = 'x'
		}
	}
}
