package main

import "fmt"

func main() {
	q, r, err := divideClient(100, 0)
	if err != nil {
		fmt.Println("do something else, err :", err)
		return
	}
	fmt.Printf("quotient = %d and remainder = %d\n", q, r)
}

func divideClient(x, y int) (quotient, remainder int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	quotient, remainder = divide(x, y)
	return
}

// 3rd party library
func divide(x, y int) (quotient, remainder int) {
	quotient, remainder = x/y, x%y
	return
}
