package main

import (
	"fmt"
	"runtime"
	"time"
)

func a() {
	for i := 1; i < 10; i++ {
		fmt.Print("A:\t", i)
	}

}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Print("B:\t", i)
	}
}

func main() {
	runtime.GOMAXPROCS(2)
	go a()
	go b()
	time.Sleep(time.Second)
}
