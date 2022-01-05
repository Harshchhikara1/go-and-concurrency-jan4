package main

import (
	"fmt"
)

func binarySearch(arr []int, num int, size int) int {
	var low, high int
	low = 0
	high = size - 1
	idx := -1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == num {
			idx = mid
			break
		} else if arr[mid] < num {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return idx
}
func main() {
	var n int
	fmt.Printf("Enter the size of array\n")
	fmt.Scanf("%d\n", &n)
	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Enter %d ", i+1)
		fmt.Println("th element of the array ")
		fmt.Scanf("%d\n", &arr[i])
	}
	fmt.Printf("Enter the element you want to search in array\n")
	var num int
	fmt.Scanf("%d", &num)
	idx := binarySearch(arr, num, n)
	if idx > -1 {
		fmt.Println("Item found at idx -> ", idx)
	} else {
		fmt.Println("Item could not be found")
	}
}
