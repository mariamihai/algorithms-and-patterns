package main

import (
	"fmt"
	"math/rand"
)

// https://www.youtube.com/watch?v=wELNUHb3kuA

func repeatFunc[T any, K any](done <-chan K, fn func() T) <-chan T {
	stream := make(chan T)

	go func() {
		defer close(stream)

		for {
			select {
			case <-done:
				return
			case stream <- fn():
			}
		}
	}()

	return stream
}

func main() {
	done := make(chan int)
	defer close(done)

	randomNumberFetcher := func() int { return rand.Intn(100000) }

	for rando := range repeatFunc(done, randomNumberFetcher) {
		fmt.Println(rando)
	}
}
