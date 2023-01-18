/*
	Channel
		Data type to facility communication between goroutines

	Declare
		var <var_name> chan <data_type>
		ex:
			var ch chan int

	Initialize
		ch = make(chan int)

	Operations
		Using <- (channel operator)

		SEND operation
			ch <- 100

		RECEIVE operation
			data := <-ch
*/

package main

import (
	"fmt"
)

//share memory by communicating
/*
func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)
	wg.Add(1)
	go add(100, 200, wg, ch)
	result := <-ch
	wg.Wait()
	fmt.Println(result)
}

func add(x, y int, wg *sync.WaitGroup, ch chan int) {
	result := x + y
	ch <- result
	wg.Done()
}
*/

func main() {
	ch := make(chan int)
	go add(100, 200, ch)
	result := <-ch
	fmt.Println(result)
}

func add(x, y int, ch chan int) {
	result := x + y
	ch <- result
}
