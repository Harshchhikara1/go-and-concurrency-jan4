package main

import "fmt" ///for _,v:=range[x]

func main() {
	var x [5]float64 {
	x[0]=96
	x[1]=76
	x[2]=80
	x[3]=

	}

	var total float64 = 0
	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	fmt.Println(total / float64(len(x)))
}
