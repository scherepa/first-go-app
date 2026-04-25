package practice

import (
	"fmt"
	"time"
)

func chanCommunicate() {

	ch := make(chan string)
	//	ch <- "lala" // send won't work as no reciever at this stage

	go func() { // inside go routine it is possible
		// Send data to the channel but what is the channel name?
		ch <- "callout:{\"id\":1}"
		// up to the end of this funcction the furthier code is blocked
	}()

	val := <-ch // receive
	fmt.Printf("got %s\n", val)

	fmt.Println("-----")

	// Send multiple
	const count = 3
	go func() {
		for i := range count {
			fmt.Printf("sending %d\n", i+1)
			ch <- fmt.Sprintf("callout:{\"id\":%d}", i+1)
			time.Sleep(time.Second)
		} // and again upto the end of the block the code is blocked
	}()

	for range count {
		val := <-ch
		fmt.Printf("received %s\n", val)
	}

	fmt.Println("---")

	// close to signal we're done
	go func() {
		for i := range count {
			fmt.Printf("sending <- %d\n", i+1)
			ch <- fmt.Sprintf("callout:{\"id\":%d}", i+1)
			time.Sleep(time.Second)
		}
		defer close(ch)
	}()

	for i := range ch { // range of channels
		fmt.Printf("received %s\n", i)
	}
	fmt.Println("--- Buffer")
	ch2 := make(chan int, 1) // buffered channel
	// buffered channel will not block if we send one value the second will...
	// what if we send via buffer csv row?
	ch2 <- 19
	val2 := <-ch2
	fmt.Println(val2)
	close(ch2)
	fmt.Println("--- Buffer END")
}
