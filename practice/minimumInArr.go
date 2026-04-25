package practice

import "fmt"

func Min[T Ordered](values []T) (T, error) {
	if len(values) == 0 {
		// zero of type T generic
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
