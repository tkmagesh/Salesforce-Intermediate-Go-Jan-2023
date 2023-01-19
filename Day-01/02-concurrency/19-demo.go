/* Using channels for signaling */

package main

import (
	"fmt"
	"time"
)

func main() {
	stopCh := make(chan bool)
	fibCh := genFib(stopCh)

	fmt.Println("Hit ENTER to stop...")
	go func() {
		fmt.Scanln()
		stopCh <- true
	}()

	for fibNo := range fibCh {
		fmt.Println(fibNo)
	}
	fmt.Println("Done")
}

func genFib(stop chan bool) <-chan int {
	ch := make(chan int)
	go func() {
		x, y := 0, 1
	LOOP:
		for {
			select {
			case <-stop:
				break LOOP
			default:
				ch <- x
				time.Sleep(500 * time.Millisecond)
				x, y = y, x+y
			}
		}
		close(ch)
	}()
	return ch
}
