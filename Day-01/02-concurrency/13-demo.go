/* sending and receiving multiple values using a channel */

package main

import (
	"fmt"
	"time"
)

//consumer
func main() {
	ch := make(chan int)
	go genNos(ch)
	for {
		if data, isOpen := <-ch; isOpen {
			fmt.Println(data)
			continue
		}
		break
	}
}

//producer
func genNos(ch chan int) {
	count := 5
	for i := 1; i <= count; i++ {
		time.Sleep(500 * time.Millisecond)
		ch <- i * 10
	}
	close(ch)
}
