package main

import "fmt"

func main() {
	x := []int{
		48, 96, 9, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 17,
	}
	MIN := x[0]
	for _, num := range x {
		if MIN > num {
			MIN = num
		}
	}
	fmt.Println(MIN)
}
