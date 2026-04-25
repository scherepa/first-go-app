package practice

import (
	"fmt"
	"svetlana/first-app/util"
	"sync"
	"time"
)

func printResCall(url string) {
	ctype, err := util.CallURLGetHeader(url)
	fmt.Printf("%#v\t|%#v\t|\t%#v\n", url, ctype, err)
}

type result struct {
	url   string
	ctype string
	err   error
}

func concurTryDiferent(url string, ch chan result) {
	res := result{url: url}
	ctype, err := util.CallURLGetHeader(url)
	res.ctype = ctype
	res.err = err
	ch <- res
}

func siteSerial2(urls []string) {
	ch2 := make(chan result)
	for _, url := range urls {
		go concurTryDiferent(url, ch2)
	}
	fmt.Print("---\nReceived From Chanel different\n---\n")
	// not really understandable the order
	for i := range len(urls) {
		val := <-ch2
		//printResCall(url)
		fmt.Printf("%d\t|%#v\n", i, val)
	}
	fmt.Println("---")
}

func siteSerial(urls []string) {
	ch := make(chan string)
	go func() {
		for _, url := range urls {
			ctype, err := util.CallURLGetHeader(url)
			ch <- fmt.Sprintf("\t%#v\t|%#v\t|%#v", url, ctype, err)
		}
		close(ch)
	}()
	fmt.Print("---\nReceived From Chanel\n---\n")
	fmt.Println("NUM\t|\t\t\t\tURL\t\t\t|\t\tContent Type\t\t|Error")
	//FILO
	for i := range len(urls) {
		val := <-ch
		//printResCall(url)
		fmt.Printf("%d\t|%s\n", i, val)
	}
	fmt.Println("---")
}

func siteRegular(urls []string) {
	fmt.Println("\t\t\t\tURL\t\t\t|\t\tContent Type\t\t|Error")
	for _, url := range urls {
		printResCall(url)
	}
	fmt.Println("---")
}

func sitesConcurrent(urls []string) {
	var wg sync.WaitGroup
	fmt.Println("\t\t\tURL\t\t\t|\t\tContent Type\t\t|\tError")
	// regular order
	for _, url := range urls {
		wg.Add(1)             // mostly waiting groups communicate via channels
		go func(url string) { // goroutine function
			printResCall(url)
			wg.Done()
		}(url)
	}
	// wait for all to settle
	wg.Wait()
	fmt.Println("---")
}

func RegularVsConcurrent() { // seems to be working like pool of http
	// but main idea more close to laravel queues which run simulteniously
	// as we use same arr to send... sites remember requests and next call will be faster anyway but obviously concurrent is faster
	urls := []string{
		"https://www.google.com/search?q=free+rest+api",
		"https://jsonplaceholder.typicode.com/posts",
		"https://golang.org",
		"https://api.github.com", // if we do not take smth very long to answer do not see much difference
	}

	// for download file urls were forbidden so stopped
	fmt.Println("---regular")
	start := time.Now()
	siteRegular(urls)
	fmt.Println(time.Since(start), "\tregular 1 by one")
	//----
	fmt.Println("---go outside for")
	start = time.Now()
	siteSerial(urls)
	fmt.Println(time.Since(start), "\tuse channel go outside")
	//---
	fmt.Println("---go wg")
	start = time.Now()
	sitesConcurrent(urls)
	fmt.Println(time.Since(start), "\tconcurrent wg")
	//---
	fmt.Println("---Start go inside for")
	start = time.Now()
	siteSerial2(urls)
	fmt.Println(time.Since(start), "\tconcurrent?")
}
