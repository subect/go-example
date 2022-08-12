package main

import "fmt"

func main() {
	slice := make([]string, 0)
	slice = append(slice, "3", "4", "5", "6")
	fmt.Println(len(slice))
	slice2 := make([]string, len(slice))
	copy(slice2, slice)
	fmt.Println(slice)
	fmt.Printf("slice:%p\n", &slice)
	fmt.Println(slice2)
	fmt.Printf("slice2:%p\n", &slice2)

}
