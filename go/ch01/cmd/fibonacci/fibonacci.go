package main

import (
	f "ch01/pkg/fibonacci"
	"fmt"
	"time"
)

func main() {
	var fibonacci = []uint{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610}

	for n := uint(0); n < 16; n++ {
		fmt.Printf("Fib3(%d) = %d, Fib2(%d) = %d, Fib(%d) = %d, actual: %d\n", n, f.Fib3(n), n, f.Fib2(n), n, f.Fib(n), fibonacci[n])
	}

	start_time := time.Now()
	_ = f.Fib(100)
	fib1_duration := time.Since(start_time)

	start_time = time.Now()
	_ = f.Fib2(100)
	fib2_duration := time.Since(start_time)

	start_time = time.Now()
	_ = f.Fib3(100)
	fib3_duration := time.Since(start_time)

	fmt.Printf("Durations: fib(100): %d, fib2(100): %d, fib3(100): %d\n", fib1_duration, fib2_duration, fib3_duration)
}
