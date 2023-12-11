package main

import (
	"fmt"
	"sync"
	"time"
)

// https://www.youtube.com/watch?v=c6DH-1nffTI

func countDown(n int) {
	for n >= 0 {
		fmt.Println(n)
		n--
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		countDown(5)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		countDown(10)
		wg.Done()
	}()

	wg.Wait()
}
