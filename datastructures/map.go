package main

import (
	"fmt"
)

func main() {
	x := make(map[int]int)
	x[1] = 10
	x[2] = 20
	fmt.Println(x[1])
	fmt.Println(x[2])
}

// func main() {
// 	x := make(map[int]int)
// 	x[1] = 10
// 	fmt.Println(x[1])
// }
