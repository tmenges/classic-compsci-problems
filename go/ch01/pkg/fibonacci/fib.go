package fibonacci

func Fib(n uint) uint {
	if n == 0 {
		return 0
	}

	if n == 1 || n == 2 {
		return 1
	}

	return Fib(n-2) + Fib(n-1)
}

func Fib2(n uint) uint {
	var cache = make(map[uint]uint)

	var fib_helper func(n uint) uint

	fib_helper = func(n uint) uint {
		switch n {
		case 0:
			cache[0] = 0
			return 0
		case 1:
			cache[1] = 1
			return 1
		default:
			n2, n2_exists := cache[n-2]
			if !n2_exists {
				n2 = fib_helper(n - 2)
			}
			n1, n1_exists := cache[n-1]
			if !n1_exists {
				n1 = fib_helper(n - 1)
			}

			cache[n] = n1 + n2
			return cache[n]
		}

	}
	return fib_helper(n)
}

func Fib3(n uint) uint {

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	var x uint = 0
	var x1 uint = 1
	var x2 uint = 0

	for i := uint(2); i <= n; i++ {
		x = x1 + x2
		x2 = x1
		x1 = x
	}

	return x
}
