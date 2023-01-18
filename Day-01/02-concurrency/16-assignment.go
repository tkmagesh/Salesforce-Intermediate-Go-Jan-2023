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
		time.Sleep(5 * time.Second)
		ch1 <- 100
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- 200
	}()

	fmt.Println(<-ch1)
	fmt.Println(<-ch2)
}
