package main

import (
	"fmt"
)

func power(x float64, a int) float64 {
	if a == 0 {
		return 1
	}
	if a == 1 {
		return x
	}
	res := x
	for i:= 1; i<a; i++ {
		res = res * x
	}
	return res
}

func main() {
	n := 1000
	e := power((1.0 + 1.0/float64(n)),n)
	fmt.Println(e)
}
