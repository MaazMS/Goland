package main

import (
	"fmt"
)

func main() {

	fmt.Println("Array example")
	var no [5]int
	fmt.Println(no)
	no[3] = 100
	fmt.Println(no)
	fmt.Println(len(no))
}
