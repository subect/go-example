package main

import "fmt"

func main() {
	mapCASE := make(map[string]string)
	mapCASE["1"] = "1"
	mapCASE["2"] = "2"
	mapCASE["3"] = "3"
	fmt.Println(len(mapCASE))
	fmt.Println(mapCASE)
	if _, ok := mapCASE["1"]; ok {
		fmt.Println(ok)
	}
	delete(mapCASE, "1")
	fmt.Println(mapCASE)
	fmt.Println("111111")
}
