// package main

// import "fmt"

// func main() {
// 	slice1 := []int{1, 2, 3}
// 	slice2 := append(slice1, 4, 5) // append
// 	fmt.Println(slice1, slice2)
// }

package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3} //copy
	slice2 := make([]int, 3)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
}
