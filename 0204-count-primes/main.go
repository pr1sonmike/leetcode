package _204_count_primes

import "math"

func countPrimes(n int) int {
	count := 0
	for i := 0; i < n; i++ {
		if isPrime(i) {
			count++
		}
	}

	return count
}

func isPrime(n int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(n)))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return n > 1
}
