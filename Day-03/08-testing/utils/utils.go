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
