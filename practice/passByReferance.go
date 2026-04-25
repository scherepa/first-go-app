package practice

import (
	"fmt"
)

func PlayAndPrint() {
	arrForUpdate := []int{1, 2, 3}
	fmt.Printf("before change Arr is: %v\n", arrForUpdate)
	DoubleArrVaueAt(arrForUpdate, 2)
	fmt.Printf("after change Arr is: %v\n", arrForUpdate)
	changedVal := 5
	fmt.Println("passing by value")
	DoubleIntNotChangingVal(changedVal)
	fmt.Printf("premitive int with value 5\nafter double by value applied:\n\tchangedVal is: %v\n", changedVal)
	fmt.Println("passing address")
	// passing address
	DoubleIntChangingVal(&changedVal)
	fmt.Printf("premitive int with value 5\nafter double by address applied:\n\tchangedVal is: %v\n", changedVal)
	// taking int
	lala := arrForUpdate[2]
	// passing address of new int lala and yes it is an integer
	// but we need to pass address to actually change it
	// as otherwise only copy is sent
	fmt.Println("passing address")
	DoubleIntChangingVal(&lala)
	fmt.Printf("Arr is: %v and lala := arrForUpdate[2]\npassed by referance &lala is: %v\n", arrForUpdate, lala)
	// lala as expected is 12 so lets pass it to double not by referance
	fmt.Println("passing value")
	DoubleIntNotChangingVal(lala)
	fmt.Printf("And now\n\tArr is: %v - same as it was\n\tand lala := arrForUpdate[2]\n\tpassed by value for double is: %v - just as it was\n", arrForUpdate, lala)
	//and what will be the arr if we do it like so?
	fmt.Println("passing address")
	DoubleIntChangingVal(&arrForUpdate[2])
	fmt.Printf("Lastly:\nAnd now after passing &arrForUpdate[2]\n\tArr is: %v - changed as expected\n\tlala is: %v - just as it was as this var has another address\n", arrForUpdate, lala)

}
