/* Anonymous functions */

package main

import "fmt"

func main() {

	func() {
		fmt.Println("fn invoked")
	}()

	func(x, y int) {
		result := x + y
		fmt.Println("Add Result :", result)
	}(100, 200)

	result := func(x, y int) int {
		return x + y
	}(100, 200)
	fmt.Println(result)

	result = func(x, y int) int {
		return x - y
	}(100, 200)
	fmt.Println(result)
}
