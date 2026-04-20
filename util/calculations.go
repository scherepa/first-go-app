package util

import (
	"fmt"
	"math"
)

func SqrtCalc(val float64) (n float64, err error) {
	if val < 0 {
		return 0.0, fmt.Errorf("sqrt can not be negative, but value passed is %f", val)
	}
	return math.Sqrt(val), nil
}
