/* Higher Order Functions - Functions as values to variables */

package main

import "fmt"

func main() {
	var fn func()
	fn = func() {
		fmt.Println("fn invoked")
	}
	fn()

	var addProcessor func(int, int)
	addProcessor = func(x, y int) {
		result := x + y
		fmt.Println("Add Result :", result)
	}
	addProcessor(100, 200)

	var mathOperation func(int, int) int
	mathOperation = func(x, y int) int {
		return x + y
	}
	fmt.Println(mathOperation(100, 200))

	mathOperation = func(x, y int) int {
		return x - y
	}
	fmt.Println(mathOperation(100, 200))
}
