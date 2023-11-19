package step1

func FibonacciRecursive(n int) int {
	if n <= 2 {
		return 1
	}

	return FibonacciRecursive(n-1) + FibonacciRecursive(n-2)
}

func FibonacciIterative(n int) int {
	f1 := 1
	f2 := 1
	for i := 2; i < n; i++ {
		f1, f2 = f2, f1+f2
	}

	return f2
}
