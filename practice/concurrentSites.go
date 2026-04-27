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

func concurrentTryDiferent(url string, ch chan customResp) {
	res := customResp{url: url}
	ctype, err := util.CallURLGetHeader(url)
	res.contentType = ctype
	res.err = err
	ch <- res
}

func siteSerial2(urls []string) {
	ch2 := make(chan customResp)
	for _, url := range urls {
		go concurrentTryDiferent(url, ch2)
	}
	fmt.Print("---\nReceived From Chanel different\n---\n")
	// not really understandable the order
	// when safe to close channel ???
	for i := range urls {
		val := <-ch2
		fmt.Println(i, val.url, val.contentType, val.err)
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
		// feeling need of control when can be closed actually
		close(ch)
	}()
	fmt.Print("---\nReceived From Chanel\n---\n")
	fmt.Println("NUM\t|\t\t\t\tURL\t\t\t|\t\tContent Type\t\t|Error")
	//as finishes...
	for i := 0; i < len(urls); i++ {
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

type customResp struct {
	url         string
	contentType string
	err         error
}

func mostControled(urls []string) {
	headers := make(chan customResp)
	var wg sync.WaitGroup
	fmt.Println("URLS COUNT", len(urls))
	fmt.Println("\t\t\tURL\t\t\t|\t\tContent Type\t\t|\tError")
	// regular order
	for _, url := range urls {
		wg.Add(1)             // mostly waiting groups communicate via channels
		go func(url string) { // goroutine function
			defer wg.Done()
			ctype, err := util.CallURLGetHeader(url)
			headers <- customResp{url, fmt.Sprintf("\t|%#v\t|", ctype), err}
		}(url)
	}
	// wait for all to settle
	go func() {
		wg.Wait()
		close(headers)
	}()
	//as finishes...
	i := 0
	for val := range headers {
		fmt.Println(i, val.url, val.contentType, val.err)
		i++
	}
	fmt.Println("---")
}

func RegularVsConcurrent() { // seems to be working like pool of http
	// but main idea more close to laravel queues which run simulteniously
	// as we use same arr to send... sites remember requests and next call will be faster anyway but obviously concurrent is faster
	urls := []string{
		"https://www.google.com/search?q=rest+api",
		"https://jsonplaceholder.typicode.com/posts/5",
		"https://jsonplaceholder.typicode.com/todos/6",
		"https://api.github.com/users/scherepa", // if we do not take smth very long to answer do not see much difference
	}

	// for download file urls were forbidden so stopped
	//fmt.Println("---regular")
	start := time.Now()
	// way too long
	siteRegular(urls)
	fmt.Println(time.Since(start), "\tregular 1 by one")
	//----
	// bad
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

	//---
	fmt.Println("---go wg and channels")
	start = time.Now()
	mostControled(urls)
	fmt.Println(time.Since(start), "\tconcurrent wg and channels controlled")
}
