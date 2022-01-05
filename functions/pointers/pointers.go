package main

import (
	"fmt"
)

// func zero(x int) {
// 	x = 0
// }
// func main() {
// 	x := 5
// 	zero(x)
// 	fmt.Println(x) // x is still 5
// }

func zero(xPtr *int) {
	*xPtr = 0
}
func main() {
	x := 5
	zero(&x)
	fmt.Println(x) // x is 0
}

// func one(xPtr *int) {
// 	*xPtr = 1
//   }
//   func main() {
// 	xPtr := new(int)
// 	one(xPtr)
// 	fmt.Println(*xPtr) // x is 1
//   }
// go has a automatic garbage collection so we don't have to worry about new
