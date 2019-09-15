package main

import "fmt"

func Add(a int32, b int32) int32 {
	return a + b
}

func main() {
	fmt.Println("Add result is {}", Add(1, 2))
}
