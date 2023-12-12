package main

import (
	f "ch01/pkg/fibonacci"
	"fmt"
)

func main() {
	var fibonacci = []uint{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610}

	for n := uint(0); n < 16; n++ {
		fmt.Printf("Fib2(%d) = %d, Fib(%d) = %d, actual: %d\n", n, f.Fib2(n), n, f.Fib(n), fibonacci[n])
	}

}
