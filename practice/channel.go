package practice

import (
	"fmt"
	"time"
)

func readSingleFromChan(ch chan string) {
	//	ch <- "lala" // send won't work as no reciever at this stage
	// inside go routine it is possible
	// Send data to the channel but what is the channel name?
	go func() {
		ch <- "lala"
	}()
	val := <-ch // receive
	fmt.Printf("got %s\n", val)
}

func putManyToChannel(ch chan string, count int) {
	for i := 0; i < count; i++ {
		fmt.Printf("sending %d\n", i+1)

		ch <- fmt.Sprintf("callout:{\"id\":%d}", i+1)
		time.Sleep(time.Second)

	}

}

func useBuffer() {
	fmt.Println("--- Buffer")
	ch2 := make(chan int, 1) // buffered channel
	// buffered channel will not block if we send one value the second will...
	// what if we send via buffer csv row? it's seems to be bad idea it is order of completion...
	ch2 <- 19
	val2 := <-ch2 // as it was buffered it is safe to read from
	fmt.Println(val2)
	close(ch2)
	fmt.Println("--- Buffer END")
}

func chanCommunicate() {
	ch := make(chan string)
	readSingleFromChan(ch)
	fmt.Println("---")
	// Send multiple
	const count = 3
	// better to use different channel and defer close channel inside go blocks but here learning
	// or go func() {} block
	go putManyToChannel(ch, count)
	for i := 0; i < count; i++ {
		val := <-ch
		fmt.Printf("received %s\n", val)
	}
	fmt.Println("---")
	// close to signal we're done
	go func() {
		for i := 0; i < count; i++ {
			fmt.Printf("sending <- %d\n", i+1)
			ch <- fmt.Sprintf("callout2:{\"id\":%d}", i+1)
			time.Sleep(time.Second)
		}
		defer close(ch)
	}()
	for i := range ch { // range of channels
		fmt.Printf("received %s\n", i)
	}
	useBuffer()
}
