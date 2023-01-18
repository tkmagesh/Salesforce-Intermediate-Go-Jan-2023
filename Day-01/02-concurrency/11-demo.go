/* sending and receiving multiple values using a channel */

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go genNos(ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func genNos(ch chan int) {
	time.Sleep(500 * time.Millisecond)
	ch <- 10
	time.Sleep(500 * time.Millisecond)
	ch <- 20
	time.Sleep(500 * time.Millisecond)
	ch <- 30
	time.Sleep(500 * time.Millisecond)
	ch <- 40
	time.Sleep(500 * time.Millisecond)
	ch <- 50
}
