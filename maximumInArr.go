package main

import (
	"fmt"
)

func FindMaxInArr() {
	arr := []int{1, 22, 5, 47, 3}
	max := arr[0]

	for _, v := range arr[1:] {
		if v > max {
			max = v
		}
	}
	fmt.Printf("The max in arr:%#v, is: %d and its type is %T\n", arr, max, max)
}
