/*
	Modify the program so that it prints
		200 (this is produced in 2 secs)
		100 (this is produced in 5 secs)
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		for i := 1; ; i++ {
			time.Sleep(500 * time.Millisecond)
			ch1 <- i * 2
		}
	}()

	go func() {
		for i := 1; ; i++ {
			time.Sleep(1 * time.Second)
			ch2 <- (i * 2) + 1
		}
	}()

	for {
		select {
		case d1 := <-ch1:
			fmt.Println(d1)
		case d2 := <-ch2:
			fmt.Println(d2)
		}
	}
}
