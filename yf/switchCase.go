package main

import (
	"fmt"
	"time"
)

func main() {
	switch time.Now().Weekday() {
	case time.Sunday, time.Saturday:
		fmt.Println("今天是星期六")
	default:
		fmt.Println("今天是工作日")
	}
}
