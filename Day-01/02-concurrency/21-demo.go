/* Using channels for signaling */

package main

import (
	"fmt"
	"time"
)

func main() {

	fibCh := genFib()
	for fibNo := range fibCh {
		fmt.Println(fibNo)
	}
	fmt.Println("Done")
}

func genFib() <-chan int {
	ch := make(chan int)
	stop := time.After(10 * time.Second)
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

//use time.After() instead
/*
func timeout(d time.Duration) <-chan time.Time {
	timeoutCh := make(chan time.Time)
	go func() {
		time.Sleep(d)
		timeoutCh <- time.Now()
	}()
	return timeoutCh
} */
