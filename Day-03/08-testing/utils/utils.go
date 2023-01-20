package utils

import "time"

func IsPrime(no int) bool {
	time.Sleep(2 * time.Second)
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

func GetPrimes(start, end int) []int {
	// var primes []int
	primes := make([]int, 0, 20)
	for i := start; i <= end; i++ {
		if IsPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}
