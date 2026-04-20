package main

import (
	"fmt"
)

func PlayAndPrint() {
	arrForUpdate := []int{1, 2, 3}
	fmt.Printf("before change Arr is: %v\n", arrForUpdate)
	DoubleArrVaueAt(arrForUpdate, 2)
	fmt.Printf("after change Arr is: %v\n", arrForUpdate)
	changedVal := 5
	// passing address
	DoubleIntChangingVal(&changedVal)
	fmt.Printf("premitive int after double by address applied:\n changedVal is: %v\n", changedVal)
	// taking int
	lala := arrForUpdate[2]
	// passing address of new int lala and yes it is an integer
	// but we need to pass address to actually change it
	// as otherwise only copy is sent
	DoubleIntChangingVal(&lala)
	fmt.Printf("Arr is: %v and lala is: %v\n", arrForUpdate, lala)
	// lala as expected is 12 so lets pass it double not by referance
	DoubleIntNotChangingVal(lala)
	fmt.Printf("And now Arr is: %v - same as it was\n and lala is: %v - just as it was\n", arrForUpdate, lala)
	//and what will be the arr if we do it like so?
	DoubleIntChangingVal(&arrForUpdate[2])
	fmt.Printf("Lastly:\nAnd now Arr is: %v - same as it was?\n and lala is: %v - just as it was\n", arrForUpdate, lala)

}
