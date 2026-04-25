package practice_test

import (
	"svetlana/first-app/practice"
	"testing"
)

var arrToTest []int = []int{1, 5, 7}

func TestDoubleArrVaueAt(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		values []int
		i      int
	}{
		{name: "test double updates value on i=2",
			values: arrToTest, i: 2},
		{name: "test double updates value on i=1",
			values: arrToTest, i: 1},
		{name: "test double updates value on i=1 after previous change",
			values: arrToTest, i: 1},
		{name: "test double updates value on i=1 in another arr",
			values: []int{5, 7, 8}, i: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lala := tt.values[tt.i]
			practice.DoubleArrVaueAt(tt.values, tt.i)
			if tt.values[tt.i] != lala*2 {
				t.Fatalf("Expected change %d to %d got %d", lala, lala*2, tt.values[tt.i])
			}
		})
	}
}
