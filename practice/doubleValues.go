package practice

// pay attention nothing returned as this is by referance
func DoubleArrVaueAt(values []int, i int) {
	values[i] *= 2
}

// no need to return anything we go to address and change value
func DoubleIntChangingVal(value *int) {
	*value *= 2
}

// no need to return anything
func DoubleIntNotChangingVal(value int) {
	value *= 2
}
