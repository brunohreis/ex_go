package main

import (
	"fmt"
	"image/gif"
	"math/rand/v2"
	"time"
)

func main() {
	anim := gif.GIF{LoopCount: nframes}
	freq := rand.Float64() * 3.0
	t := 0.0
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}
func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}
func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
