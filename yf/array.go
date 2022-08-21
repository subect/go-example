package main

import "fmt"

func main() {
	arr2 := [...]int{1, 2, 3, 4, 5}
	printArr(&arr2)
	fmt.Println(arr2)
}

func printArr(arr *[5]int) {
	arr[0] = 10
	for _, v := range arr {
		fmt.Println(v)
	}
}
