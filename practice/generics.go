package practice

import (
	"fmt"
)

type Ordered interface {
	int | float64 | string
}

func printSlice[T Ordered](s string, values []T) {
	fmt.Println("In Slice of ", s)
	fmt.Println("|\tMIN", "\t|\t", "ERROR", "\t|")
	res, err := Min(values)
	fmt.Println("|\t", res, "\t|\t", err, "\t|")
}

func RunGenerics() {
	printSlice("([]float64{2, 1, 3})", []float64{2, 1, 3})
	fmt.Println("---")
	printSlice("([]string{\"B\", \"A\", \"C\"})", []string{"B", "A", "C"})
}
