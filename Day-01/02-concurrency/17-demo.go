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
	ch3 := make(chan int)

	go func() {
		time.Sleep(5 * time.Second)
		ch1 <- 100
	}()

	go func() {
		time.Sleep(3 * time.Second)
		d3 := <-ch3
		fmt.Println("ch3 data :", d3)
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- 200
	}()

	for i := 0; i < 3; i++ {
		select {
		case d1 := <-ch1:
			fmt.Println(d1)
		case d2 := <-ch2:
			fmt.Println(d2)
		case ch3 <- 300:
			fmt.Println("data sent to ch3")
			/*
				default:
						fmt.Println("no channel operations were successful")
			*/
		}
	}
}
