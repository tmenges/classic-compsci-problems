package fibonacci

func fib(n uint) uint {
	if n == 0 {
		return 0
	}

	if n == 1 || n == 2 {
		return 1
	}

	return fib(n-2) + fib(n-1)
}

func fib2(n uint) uint {
	var cache []uint = make([]uint, 0, n)

	var fib_helper func(n uint) uint

	fib_helper = func(n uint) uint {
		switch n {
		case 0:
			cache = append(cache, 0)
			return 0
		case 1:
			cache = append(cache, 1)
			return 1
		default:
			var n2 uint
			if n-2 < uint(len(cache)) {
				n2 = cache[n-2]
			} else {
				n2 = fib_helper(n - 2)
			}
			var n1 uint
			if n-1 < uint(len(cache)) {
				n1 = cache[n-1]
			} else {
				n1 = fib_helper(n - 1)
			}

			cache = append(cache, n1+n2)
			return cache[n]
		}

	}
	return fib_helper(n)
}
