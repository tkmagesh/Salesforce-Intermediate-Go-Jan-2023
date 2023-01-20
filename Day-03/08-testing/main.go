package main

import "fmt"

func main() {
	/*
		var fn func()
		fn = func() {
			fmt.Println("fn invoked")
		}
		fn()
	*/

	var fns []func()
	for i := 0; i < 5; i++ {
		fns = append(fns, func() {
			fmt.Println("fn invoked")
		})
	}

	for _, fn := range fns {
		fn()
	}
}
