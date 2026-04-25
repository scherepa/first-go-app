package practice

import (
	"fmt"
)

type Ordered interface {
	int | float64 | string
}

func min[T Ordered](values []T) (T, error) {
	if len(values) == 0 {
		var zero T
		return zero, fmt.Errorf("Slice is empty. Minimum can not be found")
	}
	min := values[0]
	for _, val := range values[1:] {
		if val < min {
			min = val
		}
	}
	return min, nil
}

func printSlice[T Ordered](s string, values []T) {
	fmt.Println("In Slice of ", s)
	fmt.Println("|\tMIN", "\t|\t", "ERROR", "\t|")
	res, err := min(values)
	fmt.Println("|\t", res, "\t|\t", err, "\t|")
}

func RunGenerics() {
	printSlice("([]float64{2, 1, 3})", []float64{2, 1, 3})
	fmt.Println("---")
	printSlice("([]string{\"B\", \"A\", \"C\"})", []string{"B", "A", "C"})
}
