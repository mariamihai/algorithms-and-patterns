package main

import "fmt"

// https://www.youtube.com/watch?v=c6DH-1nffTI

func fibonacci(max int, ch chan int) {
	fib := make([]int, max)
	fib[0] = 0
	fib[1] = 1

	ch <- fib[0]
	ch <- fib[1]

	for i := 2; i < max; i++ {
		fib[i] = fib[i-2] + fib[i-1]
		ch <- fib[i]
	}

	close(ch)
}

func main() {
	// ch <- x // send
	// x := <- ch // receive
	// <- ch // discard

	ch := make(chan int)

	go fibonacci(20, ch)

	for msg := range ch {
		fmt.Printf("%d\n", msg)
	}
}
